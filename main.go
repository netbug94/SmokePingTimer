package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {
	interval := getInterval() // Value returned by a getInterval function, now includes user input
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the interval in minutes: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Failed to read from stdin: %s", err)
	}

	// Convert input to an integer
	minutes, err := strconv.Atoi(input[:len(input)-1]) // Remove newline character
	if err != nil {
		log.Fatalf("Invalid input, not a number: %s", err)
	}

	// Convert minutes to Duration
	return time.Duration(minutes) * time.Minute
}
