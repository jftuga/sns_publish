/*
sns_publish.go
-John Taylor
Aug 25 2021

Command line tool to publish a message to an AWS SNS topic
*/

package main

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"strings"

	"flag"
	"fmt"
	"os"
)

const pgmName string = "sns_publish"
const pgmVersion string = "1.2.0"
const pgmUrl string = "https://github.com/jftuga/sns_publish"
const maxMsgSize int = 262144

func main() {
	msgPtr := flag.String("m", "", "The message to send to the topic subscribers; surround in quotes")
	subjectPtr := flag.String("s", "", "The message subject; surround in quotes")
	topicPtr := flag.String("t", "", "The SNS topic ARN; starts with 'arn:'")
	profilePtr := flag.String("p", "", "Profile name to use; optional")
	filenamePtr := flag.String("f", "", "Send a file (instead of a message with -m)")
	version := flag.Bool("v", false, "Output version and then exit")
	flag.Parse()

	if *version {
		fmt.Printf("%v v%v\n%v\n", pgmName, pgmVersion, pgmUrl)
		return
	}

	if (*msgPtr == "" && *filenamePtr == "") || *topicPtr == "" {
		fmt.Printf("sns_publish, v%s\n", pgmVersion)
		fmt.Println(pgmUrl)
		fmt.Println("")
		fmt.Println("You must supply a subject, message and a topic ARN.")
		fmt.Println("You can optionally pass -p to use a profile other than the default.")
		fmt.Println("You can optionally pass -f to publish a file (instead of a message with -m).")
		fmt.Println("")
		fmt.Println("Usage: sns_publish -p PROFILE -s SUBJECT [-m MESSAGE|-f FILE] -t TOPIC-ARN")
		os.Exit(1)
	}

	if len(*msgPtr) > 0 && len(*filenamePtr) > 0 {
		fmt.Println()
		fmt.Println("Invalid options: -m and -f are mutually exclusive.")
		fmt.Println()
		os.Exit(1)
	}

	slots := strings.Split(*topicPtr, ":")
	if len(slots) != 6 {
		fmt.Println()
		fmt.Printf("Invalid TOPIC-ARN, %s\n", *topicPtr)
		fmt.Println("This should contain 6 colon-delimited fields and start with 'arn:'")
		fmt.Println()
		os.Exit(1)
	}
	region := slots[3]
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region), config.WithSharedConfigProfile(*profilePtr))
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%v\n", err)
		os.Exit(255)
	}

	var message *string
	if len(*msgPtr) > 0 {
		message = msgPtr
	} else {
		info, err := os.Stat(*filenamePtr)
		if err != nil {
			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()
			os.Exit(1)
		}
		if info.Size() > 256*1024 {
			fmt.Println()
			fmt.Printf("Error: file size must be less then %d bytes\n", maxMsgSize)
			fmt.Println()
			os.Exit(1)
		}
		content, err := os.ReadFile(*filenamePtr)
		if err != nil {
			fmt.Println()
			fmt.Println(err.Error())
			fmt.Println()
			os.Exit(1)
		}
		m := string(content)
		message = &m
	}

	ctx := context.TODO()
	snsClient := sns.NewFromConfig(cfg)

	result, err := snsClient.Publish(ctx, &sns.PublishInput{
		Subject:  subjectPtr,
		Message:  message,
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
