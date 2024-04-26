package main

import (
	"log"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	interval := getInterval() // Get the interval from environment variable
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Setting up a channel to listen for interruptions (SIGTERM, SIGINT)
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Run loop
	for {
		select {
		case <-ticker.C:
			runCommand("flatpak run com.google.Chrome")
		case <-sig:
			log.Println("Shutting down...")
			return
		}
	}
}

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	cmd.Stdout = os.Stdout // Redirecting stdout
	cmd.Stderr = os.Stderr // Redirecting stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Command execution failed: %s", err)
	} else {
		log.Printf("Executed command successfully: %s", command)
	}
}

func getInterval() time.Duration {
	interval := os.Getenv("RUN_INTERVAL")
	if interval == "" {
		interval = "1m" // Default interval of 1 minute
	}
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalf("Invalid interval format: %s", err)
	}
	return d
}
