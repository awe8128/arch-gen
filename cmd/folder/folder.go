package folder

import (
	"os"

	"github.com/spf13/cobra"
)

var (
	GenerateFolderCmd = &cobra.Command{
		Use:   "folder",
		Short: "Generate folders for system design",
		Long: `Create folder structure for specific system design
for example DDD
/domain
	/[domain-name]
`,
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)

func Execute() {
	err := GenerateFolderCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
