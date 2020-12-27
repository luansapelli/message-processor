package settings

import (
	"github.com/Netflix/go-env"
	"log"
)

type Settings struct {

}

var settings Settings

func init() {
	_, err := env.UnmarshalFromEnviron(&settings)
	if err != nil {
		log.Printf("Error in environment vars: %v", err)
	}
}

func GetSettings() Settings {
	return settings
}
