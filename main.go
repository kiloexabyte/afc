package main

import (
	"log"
	"time"

	"lesiw.io/cmdio/sys"
)

func main() {
	LoadTxtFile()
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

	// Check current time in Central US
	centralTime := time.Now().In(loc)
	itsFiveOrLater := centralTime.Hour() >= 17

	if !itsFiveOrLater {
		OpenApps(rnr)
	} 

	for {
		if itsFiveOrLater {
			break
		} 
		
		Wait()

		// Send Scroll Lock key press
		loadAssembly := "[System.Reflection.Assembly]::" +
			"LoadWithPartialName('System.Windows.Forms'); "
		sendKey := "[System.Windows.Forms.SendKeys]::SendWait('%{SCROLLLOCK}')"
		cmd := loadAssembly + sendKey

		err := rnr.Run("powershell", "-Command", cmd)
        if err != nil {
            log.Printf("Error sending Scroll Lock: %v\n", err)
        }
	}
}
