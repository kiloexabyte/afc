package main

import (
	"log"
	"math/rand"
	"time"

	"lesiw.io/cmdio/sys"
)

func main() {
	var rnr = sys.Runner().WithEnv(map[string]string{
		"PKGNAME":     "cmdio",
		"CGO_ENABLED": "0",
	})

	defer rnr.Close()

	err := rnr.Run("echo", "hello from", rnr.Env("PKGNAME"))
	if err != nil {
		log.Fatal(err)
	}

	// Load Central Time location
	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Random duration between 30 to 100 seconds
		randomSeconds := rand.Intn(71) + 30 // 0-70 + 30 = 30-100
		time.Sleep(time.Duration(randomSeconds) * time.Second)

		// Check if current time in Central US is 5 PM
		centralTime := time.Now().In(loc)
		if centralTime.Hour() == 17 {
			log.Println("It's 5 PM Central Time. Exiting program.")
			return
		}

		// Send Scroll Lock key press
		err := rnr.Run("powershell", "-Command", "[System.Windows.Forms.SendKeys]::SendWait('%{SCROLLLOCK}')")
		if err != nil {
			log.Printf("Error sending Scroll Lock: %v\n", err)
		}
	}
}
