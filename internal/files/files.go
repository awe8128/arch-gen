package files

import (
	"fmt"
	"os"
	"path/filepath"
)

func Generate(content, dir, filename string) error {
	path := filepath.Join(fmt.Sprintf("./be/%s", dir), filename)
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}

	return nil
}
