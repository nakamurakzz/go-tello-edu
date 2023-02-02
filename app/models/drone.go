package models

import (
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

const (
	DefaultSpeed = 20
	WaitTime     = 5 * time.Second
)

type DroneManager struct {
	*tello.Driver
	Speed int
}

func NewDroneManager() *DroneManager {
	drone := tello.NewDriver("8889")
	droneManager := &DroneManager{
		Driver: drone,
		Speed:  DefaultSpeed,
	}

	work := func() {
		// TODO: add work
	}
	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	go robot.Start()

	time.Sleep(WaitTime)

	return droneManager
}
