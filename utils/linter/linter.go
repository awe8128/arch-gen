package linter

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// FormatAndFixImportsDir runs goimports (or gofmt fallback) on all .go files under dir.
// No return: it logs errors and continues.
func FormatAndFixImportsDir(dir string) {
	files := make([]string, 0, 64)

	err := filepath.WalkDir(dir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil {
			log.Printf("walk error: %s: %v", path, walkErr)
			return nil // keep walking
		}
		if d.IsDir() {
			// skip common dirs
			switch d.Name() {
			case ".git", "vendor", "node_modules", "dist", "build":
				return filepath.SkipDir
			}
			return nil
		}
		if strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("WalkDir failed for %s: %v", dir, err)
	}

	if len(files) == 0 {
		log.Printf("no .go files found under %s", dir)
		return
	}

	// Prefer goimports on all files at once (faster)
	if err := exec.Command("goimports", append([]string{"-w"}, files...)...).Run(); err == nil {
		return
	} else {
		log.Printf("goimports failed for dir %s: %v (falling back to gofmt)", dir, err)
	}

	// Fallback: gofmt
	if err := exec.Command("gofmt", append([]string{"-w"}, files...)...).Run(); err != nil {
		log.Printf("gofmt failed for dir %s: %v", dir, err)
	}
}
