# sns_publish
command line tool to send a message to all AWS SNS topic subscribers

# usage:
```
You must supply a subject, message and a topic ARN.

Usage: sns_publish -p PROFILE -s SUBJECT -m MESSAGE -t TOPIC-ARN
```

# adopted from:
* https://docs.aws.amazon.com/sdk-for-go/v1/developer-guide/sns-example-publish.html
* https://github.com/awsdocs/aws-doc-sdk-examples/blob/master/go/example_code/sns/SnsPublish.go

