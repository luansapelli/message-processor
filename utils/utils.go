package utils

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"golang-event-processor/settings"
	"log"
	"math/rand"
	"sync"
	"time"
)

func CreateSqsClient() sqsiface.SQSAPI {

	awsSession := session.Must(session.NewSession())

	awsConfig := &aws.Config{
		Region:   aws.String(settings.GetSettings().AwsRegion),
		Endpoint: aws.String(settings.GetSettings().QueueUrl),
	}

	return sqs.New(awsSession, awsConfig)
}

func CreateFirehoseSession() *firehose.Firehose {

	var onceFirehose sync.Once
	var firehoseClient *firehose.Firehose

	onceFirehose.Do(func() {

		rand.Seed(time.Now().UnixNano())

		instance, err := session.NewSession()
		if err != nil {
			log.Printf("Error on AWS session - %s", err.Error())
			return
		}

		_, err = instance.Config.Credentials.Get()
		if err != nil {
			log.Printf("Error to get AWS credentials - %s", err.Error())
		}

		firehoseClient = firehose.New(instance)
	})

	return firehoseClient
}
