package main

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambdaeventsources"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssqs"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

type SqsLambdaDemoStackProps struct {
	awscdk.StackProps
}

func NewSqsLambdaDemoStack(scope constructs.Construct, id string, props *SqsLambdaDemoStackProps) awscdk.Stack {
	var sprops awscdk.StackProps
	if props != nil {
		sprops = props.StackProps
	}
	stack := awscdk.NewStack(scope, &id, &sprops)

	queue := awssqs.NewQueue(stack, jsii.String("training-academy-christian-queue"), &awssqs.QueueProps{
		VisibilityTimeout: awscdk.Duration_Seconds(jsii.Number(300)),
	})

	sqsEventSource := awslambdaeventsources.NewSqsEventSource(queue, nil)

	sqsLambda := awslambda.NewFunction(stack, jsii.String("training-academy-christian-lambda"), &awslambda.FunctionProps{
		Handler: jsii.String("lambda_handler.lambda_handler"),
		// Runtime:      awslambda.Runtime_PROVIDED_AL2(),
		Runtime:      awslambda.Runtime_PYTHON_3_11(),
		Architecture: awslambda.Architecture_ARM_64(),
		Role: awsiam.NewRole(stack, jsii.String("stori-training"), &awsiam.RoleProps{
			AssumedBy: awsiam.NewServicePrincipal(jsii.String("lambda.amazonaws.com"), nil),
		}),
		Code: awslambda.AssetCode_FromAsset(jsii.String("lambda"), nil),
	})

	sqsLambda.AddEventSource(sqsEventSource)

	return stack
}

func main() {
	defer jsii.Close()

	app := awscdk.NewApp(nil)

	NewSqsLambdaDemoStack(app, "training-academy-christian-stack", &SqsLambdaDemoStackProps{
		awscdk.StackProps{
			Env: env(),
		},
	})

	app.Synth(nil)
}

// env determines the AWS environment (account+region) in which our stack is to
// be deployed. For more information see: https://docs.aws.amazon.com/cdk/latest/guide/environments.html
func env() *awscdk.Environment {
	return &awscdk.Environment{
		Account: jsii.String("832372344708"),
		Region:  jsii.String("us-east-1"),
	}
}
