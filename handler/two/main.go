package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/gino8070/aws-sam-with-sfn/utils"
)

func main() {
	lambda.Start(handler)
}

func handler(ctx context.Context, input utils.SfnOutput) (*utils.SfnOutput, error) {
	fmt.Println("start handler")

	bytes, _ := json.Marshal(&input)
	fmt.Printf("input data %v \n", string(bytes))

	out := &utils.SfnOutput{
		Msg: "two",
	}
	fmt.Println("finish handler")
	return out, nil
}
