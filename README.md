# Golang Event Processor 🚀

### Tools

* [Localstack](https://github.com/localstack/localstack)
* [AWSLocal CLI](https://github.com/localstack/awscli-local)

### Step by step

1. Copie o `.env.example` para `.env` e carregue as variáveis de ambiente
2. Suba a localstack `SERVICES=sns,sqs localstack start`
3. Execute o `create_env.sh` para criar o tópico SNS, a fila SQS e o subscribe do SQS no SNS
4. Suba um Firehose na Amazon (criei na conta pessoal da amazon)
5. Execute `go run main.go`
6. Envie uma mensagem para o SNS utilizando `awslocal sns publish --topic-arn arn:aws:sns:us-east-1:000000000000:event-dev-sns --message '{"insira":"sua mensagem"}'`