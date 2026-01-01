package main

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"

	"lesiw.io/cmdio"
)

var appDirs []string


func LoadTxtFile() {
	file, err := os.Open("appDirs.txt")
	if err != nil {
		// Use default directories if file doesn't exist
		appDirs = []string{}
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			appDirs = append(appDirs, line)
		}
	}
}

func OpenApps(rnr *cmdio.Runner) {
    for _, appDir := range appDirs {
        // Extract app name from path
        appName := strings.TrimSuffix(strings.Split(appDir, "\\")[len(strings.Split(appDir, "\\"))-1], ".exe")
        
        // Check if process is running
        cmd := exec.Command("tasklist", "/FI", "IMAGENAME eq "+appName+".exe")
        output, err := cmd.Output()
        
        if err == nil && strings.Contains(string(output), appName) {
            log.Printf("App %s is already running\n", appName)
            continue
        }
        
        // App is not running, start it
        err = rnr.Run("powershell", "-Command", "Start-Process '"+appDir+"'")
        if err != nil {
            log.Printf("Error opening app %s: %v\n", appDir, err)
        } else {
            log.Printf("Started app %s\n", appDir)
        }
    }
}

// func CloseApps(rnr *cmdio.Runner) {
// 	for _, appDir := range appDirs {
// 		err := rnr.Run("powershell", "-Command", "Stop-Process -Name '"+appDir+"' -ErrorAction SilentlyContinue")
// 		if err != nil {
// 			log.Printf("Error closing app %s: %v\n", appDir, err)
// 		}
// 	}
// }