package content

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateEntity(dir, name string) {

	content := fmt.Sprintf(
		`package %s

import "fmt"
// This file was created manually by a Go program.

func ManualFunc() {
	fmt.Println("This is a manually created function.")
}
`, name,
	)

	path := filepath.Join(dir, "test.go")
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		log.Fatalf("Failed to create file: %s: %v", "test.go", err)
	}

}
