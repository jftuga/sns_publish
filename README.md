# sns_publish
command line tool to send a message to all AWS SNS topic subscribers

You must supply a subject, message (or a file) and a topic ARN.
* You can optionally pass -p to use a profile other than the default.
* You can optionally pass -f to publish a file (instead of a message with -m).

## Usage
```
Usage: sns_publish -p PROFILE -s SUBJECT [-m MESSAGE|-f FILE] -t TOPIC-ARN

  -f string
    	Send a file (instead of a message with -m)
  -m string
    	The message to send to the topic subscribers; surround in quotes
  -p string
    	Profile name to use; optional
  -s string
    	The message subject; surround in quotes
  -t string
    	The SNS topic ARN; starts with 'arn:'
  -v	Output version and then exit
```

## Installation
* macOS: `brew update; brew install jftuga/tap/sns_publish`
* Binaries for Linux, macOS and Windows are provided in the [releases](https://github.com/jftuga/sns_publish/releases) section.
