package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gino8070/aws-sam-with-sfn/utils"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context) (*utils.SfnOutput, error) {
	fmt.Println("start handler")
	out := &utils.SfnOutput{
		Msg: "one",
	}
	fmt.Println("finish handler")
	return out, nil
}
