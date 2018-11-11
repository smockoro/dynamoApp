package main

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type MessageEvent struct {
	Recipient string `json:"recipient"`
}

type MessageResponse struct {
	Recipient string `json:"recipient"`
	Date      string `json:"Date"`
	Sender    string `json:"sender"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
}

func HandleRequest(ctx context.Context, event MessageEvent) ([]*MessageResponse, error) {
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))
	params := &dynamodb.QueryInput{
		TableName: aws.String("Message"),
		ExpressionAttributeNames: map[string]*string{
			"#Recipient": aws.String("Recipient"),
			"#Date":      aws.String("Date"),
			"#Sender":    aws.String("Sender"),
			"#Subject":   aws.String("Subject"),
			"#Message":   aws.String("Message"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":recipient": {
				S: aws.String(event.Recipient),
			},
		},
		KeyConditionExpression: aws.String("#Recipient = :recipient"),
		ProjectionExpression:   aws.String("#Recipient, #Date, #Sender, #Subject, #Message"),
		ConsistentRead:         aws.Bool(false),
		Limit:                  aws.Int64(10),
	}
	result, err := svc.Query(params)
	if err != nil {
		log.Fatal(err)
	}
	if result == nil {
		return nil, nil
	}
	resp := make([]*MessageResponse, 0)
	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &resp)
	if err != nil {
		log.Fatal(err)
	}
	return resp, nil
}

func main() {
	lambda.Start(HandleRequest)
}
