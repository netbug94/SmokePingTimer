import (
	"os/exec"
	"os"
	"time"
	"log"
)

func main() {
	// Schedule the task to run every minute
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		runCommand("sudo -n true")
	}
}

func runCommand(command string) {
	// Splitting the command into command name and arguments
	cmdParts := exec.Command("sh", "-c", command)
	cmdParts.Stdout = os.Stdout // Correctly redirecting stdout
	cmdParts.Stderr = os.Stderr // Correctly redirecting stderr

	// Start and wait for the command to finish
	if err := cmdParts.Run(); err != nil {
		log.Fatalf("Command execution failed: %s", err)
	}
}
