package files

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

func CreateFiles() {

	folders := viper.Get("domain").(map[string]any)

	for key := range folders {
		filename := fmt.Sprintf("%s.go", key)
		if err := os.WriteFile(filename, []byte("content"), 0644); err != nil {
			log.Fatalf("Failed to create file: %s", filename)
		}

	}
}
