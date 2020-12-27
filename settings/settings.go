package settings

import (
	"github.com/Netflix/go-env"
	"log"
)

type Settings struct {
	AwsRegion   string `env:"AWS_REGION"`
	QueueUrl    string `env:"QUEUE_URL"`
	WorkerCount string `env:"WORKER_COUNT"`
	StreamName  string `env:"STREAM_NAME"`
}

var settings Settings

func init() {
	_, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Printf("Error on environment vars - %s", err)
	}
}

func GetSettings() Settings {
	return settings
}
