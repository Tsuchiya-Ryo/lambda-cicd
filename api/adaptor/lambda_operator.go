package adaptor

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"lambda-cicd/api/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

type LambdaOperator struct {
	Client *lambda.Client
}

func NewLambdaOperator() *LambdaOperator {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	return &LambdaOperator{
		Client: lambda.NewFromConfig(cfg),
	}
}

func (l *LambdaOperator) Invoke(input model.InvokeInput) (*lambda.InvokeOutput, error) {
	b, err := json.Marshal(input)
	if err != nil {
		return nil, err
	}
	p := &lambda.InvokeInput{
		FunctionName: aws.String(os.Getenv("LAMBDA_FUNCTION_ARN")),
		Payload:      b,
	}
	resp, err := l.Client.Invoke(context.TODO(), p)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
