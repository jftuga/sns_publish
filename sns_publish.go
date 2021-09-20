/*
sns_publish.go
-John Taylor
Aug 25 2021

Command line tool to publish a message to an AWS SNS topic

Adapted from:
https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sns-example-publish.html

*/
package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"strings"

	"flag"
	"fmt"
	"os"
)

const pgmVersion string = "1.0.1"
const pgmUrl string = "https://github.com/jftuga/sns_publish"

func main() {
	msgPtr := flag.String("m", "", "The message to send to the subscribed users of the topic")
	subjectPtr := flag.String("s", "", "The message subject")
	topicPtr := flag.String("t", "", "The ARN of the topic to which the user subscribes")

	flag.Parse()

	if *msgPtr == "" || *topicPtr == "" {
		fmt.Printf("sns_publish, v%s\n", pgmVersion)
		fmt.Println(pgmUrl)
		fmt.Println("")
		fmt.Println("You must supply a subject, message and a topic ARN.")
		fmt.Println("Your default AWS credentials must have permission to publish messages to the given topic.")
		fmt.Println("")
		fmt.Println("Usage: sns_publish -s SUBJECT -m MESSAGE -t TOPIC-ARN")
		os.Exit(1)
	}

	slots := strings.Split(*topicPtr, ":")
	if len(slots) != 6 {
		fmt.Println()
		fmt.Printf("Invalid TOPIC-ARN, %s\n", *topicPtr)
		fmt.Println("This should contain 6 colon-delimited fields.")
		fmt.Println()
		os.Exit(1)
	}
	region := slots[3]

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		fmt.Println()
		fmt.Println(err.Error())
		os.Exit(1)
	}

	svc := sns.New(sess)
	result, err := svc.Publish(&sns.PublishInput{
		Subject:  subjectPtr,
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println()
		fmt.Println(err.Error())
		fmt.Println()
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
}
