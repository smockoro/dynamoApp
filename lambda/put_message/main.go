package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type MessageEvent struct {
	Recipient string `json:"recipient"`
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

type MessageResponse struct {
	Result *dynamodb.PutItemOutput
}

func HandleRequest(ctx context.Context, event MessageEvent) (MessageResponse, error) {
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Message"),
		Item: map[string]*dynamodb.AttributeValue{
			"Recipient": {
				S: aws.String(event.Recipient),
			},
			"Date": {
				S: aws.String(time.Now().String()),
			},
			"Sender": {
				S: aws.String(event.Sender),
			},
			"Subject": {
				S: aws.String(event.Subject),
			},
			"Message": {
				S: aws.String(event.Message),
			},
		},
	}

	resp, err := svc.PutItem(params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(resp)
	return MessageResponse{Result: resp}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
