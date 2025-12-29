package sqlc

import (
	"log"
	"os/exec"
)

func RunSQLC() {
	cmd := exec.Command("sqlc", "generate")
	cmd.Dir = "be"
	if err := cmd.Run(); err == nil {
		return
	} else {
		log.Printf("sqlc failed: %v", err)
	}
}
