package fs

import (
	"fmt"
	"os"
	"path/filepath"
)

func GenerateFile(content, dir, filename string) error {
	path := filepath.Join(fmt.Sprintf("./%s", dir), filename)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
