package main

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func makeMessage(recipient, sender, subject, message string) *dynamodb.PutItemInput {
	params := &dynamodb.PutItemInput{
		TableName: aws.String("Message"),
		Item: map[string]*dynamodb.AttributeValue{
			"Recipient": {
				S: aws.String(recipient),
			},
			"Date": {
				S: aws.String(time.Now().String()),
			},
			"Sender": {
				S: aws.String(sender),
			},
			"Subject": {
				S: aws.String(subject),
			},
			"Message": {
				S: aws.String(message),
			},
		},
	}
	return params
}

func main() {
	// Client Initialized
	svc := dynamodb.New(session.New(), aws.NewConfig().WithRegion("ap-northeast-1"))

	// Data
	recipents := []string{"David", "Alice"}
	senders := []string{"Bob", "John", "Alice"}
	subjects := []string{
		"Schedule Confirmation for June 15th",
		"Deadline for Project Budget Analysis",
		"Request for Product Documents",
		"Invitation for June 15th Conference Call",
		"Urgent Request Concerning Your Invoice",
		"Re:Schedule Confirmation for June 15th",
		"Re:Deadline for Project Budget Analysis",
		"Re:Request for Product Documents",
		"Re:Invitation for June 15th Conference Call",
		"Re:Urgent Request Concerning Your Invoice",
		"Re:Re:Schedule Confirmation for June 15th",
		"Re:Re:Deadline for Project Budget Analysis",
		"Re:Re:Request for Product Documents",
		"Re:Re:Invitation for June 15th Conference Call",
		"Re:Re:Urgent Request Concerning Your Invoice",
		"Fw:Schedule Confirmation for June 15th",
		"Fw:Deadline for Project Budget Analysis",
		"Fw:Request for Product Documents",
		"Fw:Invitation for June 15th Conference Call",
		"Fw:Urgent Request Concerning Your Invoice",
	}

	messages := []string{
		"We are happy to inform you that...",
		"We are happy to inform you that we received your payment.",
		"We are pleased to inform you that we will accept your offer. ",
		"We are sorry to inform you that...",
		"We are sorry to inform you that the item you ordered is out of stock.",
		"I regret to inform you that I would like to cancel the following order.",
		"Unfortunately, we will need to put this deal on hold.",
		"Unfortunately, there has been a change of plans.",
		"I need to ask an urgent favor of you. Would you mind sending me copies of the documents by noon?",
		"I am sorry this is very last minute, but could we meet to discuss Project A tomorrow morning?",
		"Please excuse my lack of a prompt reply as I was in a meeting. ",
		"Sorry for my late reply as I have been away on business.",
		"I was not able to reply sooner because I did not have access to the Internet.",
		"The reason I am writing this e-mail to you is because I would like to know your availability for next week to schedule a meeting.",
		"The purpose of this e-mail is to ask some questions about your current product.",
		"I am writing this e-mail to follow-up on what we discussed at the meeting.",
		"Thank you very much for your reply.",
		"Thank you very much for your help.",
		"Thank you very much for setting the agenda.",
		"I appreciate you arranging your schedule.",
	}

	// Data Import
	for _, recipent := range recipents {
		for i := 0; i < len(messages); i++ {
			params := makeMessage(recipent, senders[i%len(senders)], subjects[i], messages[i])
			_, err := svc.PutItem(params)
			if err != nil {
				panic(err)
			}

		}
	}
}
