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
	interval := getInterval() // Value returned by a getInterval function
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	// Catches interruptions
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	// Loop
	for {
		select {
		case <-ticker.C:
			runCommand("sudo -n true") // Command to be run
		case <-sig:
			log.Println("Shutting down...")
			return
		}
	}
}

func runCommand(command string) {
	cmd := exec.Command("sh", "-c", command)
	// Redirect the output and error streams to your screen:
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		log.Printf("Command execution failed: %s", err)
	} else {
		log.Printf("Executed command successfully: %s", command)
	}
}

func getInterval() time.Duration {
	const day = 24 * time.Hour // 24 hours
	defaultInterval := 7 * day // 7 x 24...

	interval := os.Getenv("RUN_INTERVAL")
	if interval == "" {
		interval = defaultInterval.String()
	}
	d, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatalf("Invalid interval format: %s", err)
	}
	return d
}
