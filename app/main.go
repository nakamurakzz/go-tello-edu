package main

import (
	"log"
	"time"

	"example/go-tello-edu/app/config"
	// "example/go-tello-edu/app/models"
	"example/go-tello-edu/app/utils"
)

const (
	WaitTime = 5 * time.Second
)

func main() {
	utils.LoggingSettings(config.Config.LogFileName)
	log.Println("test")
	// doroneManager := models.NewDroneManager()
	// doroneManager.TakeOff()
	// time.Sleep(10 * time.Second)

	// doroneManager.Land()
}
