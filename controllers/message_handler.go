package controllers

import (
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/firehose"
	"github.com/aws/aws-sdk-go/service/sqs"
	"golang-event-processor/settings"
	"log"
	"sync"
)

func (worker *Worker) processMessages(messages []*sqs.Message) {

	var waitGroup sync.WaitGroup

	numMessages := len(messages)

	waitGroup.Add(numMessages)
	log.Println("Processing messages...")

	for i := range messages {
		go func(message *sqs.Message) {

			defer waitGroup.Done()

			if err := worker.processMessage(message); err != nil {
				log.Printf("Error in message - %+v\n", *message.MessageId)
				log.Printf("Error to process message - %+v\n", err)
			} else {
				err = worker.deleteMessage(message)
				if err != nil {
					log.Printf("Error to delete message - %+v\n", err)
				}
				log.Printf("Success to process message - %+v\n", *message.MessageId)
			}
		}(messages[i])
	}
	waitGroup.Wait()
}

func (worker *Worker) processMessage(message *sqs.Message) error {

	type Message struct {
		Message string `json:"Message"`
	}

	var sqsMessage Message
	var FirehoseMessage map[string]interface{}

	if err := json.Unmarshal([]byte(*message.Body), &sqsMessage); err != nil {
		log.Printf("Error to Unmarshal - %+v\n", err)
		return err
	}

	messageBody := sqsMessage.Message

	if err := json.Unmarshal([]byte(messageBody), &FirehoseMessage); err != nil {
		log.Printf("Error to Unmarshal - %+v\n", err)
		return err
	}

	err := worker.sendMessageToFirehose(FirehoseMessage)
	if err != nil {
		log.Printf("Error to send message to Firehose - %v", err)
		return err
	}
	return nil
}

func (worker *Worker) sendMessageToFirehose(message map[string]interface{} ) error {

	messageJSON, err := json.Marshal(message)

	_, err = worker.FirehoseClient.PutRecord(&firehose.PutRecordInput{
		DeliveryStreamName: aws.String(settings.GetSettings().StreamName),
		Record:             &firehose.Record{Data: messageJSON},
	})

	if err != nil {
		return err
	}

	log.Println("Success to send message to Firehose")
	return nil
}

func (worker *Worker) deleteMessage(message *sqs.Message) error {

	params := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(worker.Config.QueueURL), // Required
		ReceiptHandle: message.ReceiptHandle,              // Required
	}
	_, err := worker.SqsClient.DeleteMessage(params)

	if err != nil {
		return err
	}
	return nil
}
