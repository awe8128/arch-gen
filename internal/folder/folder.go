package folder

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

func CreateFolders() {

	folders := viper.Get("domain").(map[string]any)

	for key := range folders {
		if err := os.Mkdir(key, 0o755); err != nil {
			log.Fatalf("Failed to create folder: %v", err)
		}
	}

}
