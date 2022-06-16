package main

import (
	"fmt"
	"os"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/dji/tello"
)

func main() {
	drone := tello.NewDriver("8889")

	work := func() {

		if err := drone.TakeOff(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to take off: %s", err)
		}

		// NOTE: When the duration time is short, Tello sometimes doesn't land.
		// In fact, when it was 5 seconds, my Tello did not land.
		gobot.After(10*time.Second, func() {
			if err := drone.Land(); err != nil {
				fmt.Fprintf(os.Stderr, "failed to land: %s", err)
			}
		})
	}

	robot := gobot.NewRobot("tello",
		[]gobot.Connection{},
		[]gobot.Device{drone},
		work,
	)

	if err := robot.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "failed to start: %s", err)
	}
}
