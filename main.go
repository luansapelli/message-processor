package main

import (
	"golang-event-processor/controllers"
)

func main() {
	go controllers.HealthCheck()
	controllers.StartPool()
}
