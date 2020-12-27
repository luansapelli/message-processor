awslocal sns create-topic --name local-sns-topic
awslocal sqs create-queue --queue-name local-sqs-queue
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:local-sns-topic --notification-endpoint http://localhost:4566/000000000000/local-sqs-queue --protocol sqs
