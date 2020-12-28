# Golang Message Processor 🚀

<p>
This service basically serves to get messages from a SQS queue, send these messages to Kinesis Data Firehose in real time and Firehose persists the messages in an S3 Bucket.
</p>

### Tools

* [Localstack](https://github.com/localstack/localstack)
* [AWSLocal CLI](https://github.com/localstack/awscli-local)

### Step by step

1. Start the localstack using the file `docker-compose-yaml` (make sure the Docker is running)
2. Execute the `create_env.sh` file to create an SNS topic, an SQS queue, and the subscription (check both were created and are running)
3. Create a Kinesis Data Firehose (at the same region as SQS/SNS) and set the destination to Amazon S3 bucket (I've created both on my personal amazon account)
4. Set the enviroment vars (you can use the `.env.example`)
5. Execute the service golang-message-processor
6. Send messages (you can send more than one message) from terminal to SNS using `awslocal sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:local-sns-topic --message '{"put":"your message"}'`
7. If everything goes correctly, you will see the following message in terminal `Success to process message`
8. Wait a few minutes (don't worry, Kinesis takes a while to process), check your Amazon S3 bucket and verify the created file, your message will be there, I promise

