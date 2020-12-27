awslocal sns create-topic --name local-event-sns
awslocal sqs create-queue --queue-name local-queue-sqs
awslocal sns subscribe --topic-arn arn:aws:sns:us-east-1:000000000000:local-event-sns --notification-endpoint http://localhost:4566/000000000000/local-queue-sqs --protocol sqs
