package controllers

import (
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/service/sqs/sqsiface"
	"golang-event-processor/utils"
	"log"
	"sync"
)

type QueueSettings struct {
	MaxNumberOfMessage int64
	QueueURL           string
	WaitTimeSecond     int64
}

type Worker struct {
	Config         *QueueSettings
	SqsClient      sqsiface.SQSAPI
	FirehoseClient *firehose.Firehose
}

func (queueSettings *QueueSettings) populateDefaultValues() {
	if queueSettings.MaxNumberOfMessage == 0 {
		queueSettings.MaxNumberOfMessage = 10
	}

	if queueSettings.WaitTimeSecond == 0 {
		queueSettings.WaitTimeSecond = 20
	}
}

func New(client sqsiface.SQSAPI, queueSettings *QueueSettings) *Worker {
	queueSettings.populateDefaultValues()

	return &Worker{
		Config:         queueSettings,
		SqsClient:      client,
		FirehoseClient: utils.CreateFirehoseSession(),
	}
}

func worker(ctx context.Context, workerConfig *QueueSettings, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	sqsClient := utils.CreateSqsClient()
	eventWorker := New(sqsClient, workerConfig)
	eventWorker.Start(ctx)
}

func (worker *Worker) Start(ctx context.Context) {

	for {

		select {

		case <-ctx.Done():
			log.Println("Worker: Stopping polling because a context kill signal was sent!")
			return

		default:

			params := &sqs.ReceiveMessageInput{
				QueueUrl:            aws.String(worker.Config.QueueURL), // Required
				MaxNumberOfMessages: aws.Int64(worker.Config.MaxNumberOfMessage),
				AttributeNames: []*string{
					aws.String("All"), // Required
				},
				WaitTimeSeconds:   aws.Int64(worker.Config.WaitTimeSecond),
				VisibilityTimeout: aws.Int64(30),
			}

			resp, err := worker.SqsClient.ReceiveMessage(params)
			if err != nil {
				log.Printf("Error to Receive Message - %v", err)
				continue
			}

			if len(resp.Messages) > 0 {
				worker.processMessages(resp.Messages)
			}
		}
	}
}
