# golang-message-processor 🧙🏽‍♂️

This service serves to get messages from a SQS queue, send these messages to Kinesis Data Firehose.

### Step by step

- Start the localstack using the file `docker-compose-yaml` (make sure the Docker is running).
- Execute the `create_env.sh` file to create an SNS topic, an SQS queue, and the subscription (check both were created and are running).
- Create a Kinesis Data Firehose (at the same region as SQS/SNS) and set the destination to Amazon S3 bucket (I've created both on my personal amazon account).
- Set the enviroment vars (you can use the `.env.example`).
- Execute the service golang-message-processor.
- Send messages (you can send more than one message) from terminal to SNS using `awslocal sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:local-sns-topic --message '{"put":"your message"}'`.
- If everything goes correctly, you will see the following message in terminal `Success to process message`.
- Wait a few minutes (don't worry, Kinesis takes a while to process), check your Amazon S3 bucket and verify the created file, your message will be there.
