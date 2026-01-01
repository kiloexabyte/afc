package main

import (
	"math/rand"
	"time"
)

func Wait() {
	// Random duration between 30 to 100 seconds
	randomSeconds := rand.Intn(71) + 30 // 0-70 + 30 = 30-100
	time.Sleep(time.Duration(randomSeconds) * time.Second)
}