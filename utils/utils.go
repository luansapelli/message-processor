package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"golang-event-processor/settings"
)

func CreateSqsClient() sqsiface.SQSAPI {

	awsSession := session.Must(session.NewSession())

	awsConfig := &aws.Config{
		Region:   aws.String(settings.GetSettings().AwsRegion),
		Endpoint: aws.String(settings.GetSettings().QueueUrl),
	}

	return sqs.New(awsSession, awsConfig)
}
