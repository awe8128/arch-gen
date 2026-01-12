package sqlc

import (
	"os/exec"
)

func RunSQLC() error {
	cmd := exec.Command("sqlc", "generate")
	cmd.Dir = "be"
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}
