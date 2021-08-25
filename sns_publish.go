// snippet-comment:[These are tags for the AWS doc team's sample catalog. Do not remove.]
// snippet-sourceauthor:[Doug-AWS]
// snippet-sourcedescription:[SnsPublish.go -m MESSAGE -t TOPIC-ARN sends MESSAGE to all the users subscribed to the SNS topic with the ARN TOPIC-ARN.]
// snippet-keyword:[Amazon Simple Notification Service]
// snippet-keyword:[Amazon SNS]
// snippet-keyword:[Publish function]
// snippet-keyword:[Go]
// snippet-sourcesyntax:[go]
// snippet-service:[sns]
// snippet-keyword:[Code Sample]
// snippet-sourcetype:[full-example]
// snippet-sourcedate:[2020-1-6]
/*
   Copyright 2010-2020 Amazon.com, Inc. or its affiliates. All Rights Reserved.

   This file is licensed under the Apache License, Version 2.0 (the "License").
   You may not use this file except in compliance with the License. A copy of
   the License is located at

    http://aws.amazon.com/apache2.0/

   This file is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
   CONDITIONS OF ANY KIND, either express or implied. See the License for the
   specific language governing permissions and limitations under the License.
*/
// snippet-start:[sns.go.publish]
package main

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"

	"flag"
	"fmt"
	"os"
)

const pgmVersion string = "1.0.0"
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

	// Initialize a session that the SDK will use to load
	// credentials from the shared credentials file. (~/.aws/credentials).
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := sns.New(sess)

	result, err := svc.Publish(&sns.PublishInput{
		Subject:  subjectPtr,
		Message:  msgPtr,
		TopicArn: topicPtr,
	})
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(*result.MessageId)
}

// snippet-end:[sns.go.publish]
