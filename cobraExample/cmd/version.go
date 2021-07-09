package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("just CLOUD MANAGER -- v3.141592653589793238462643383279502884197169399375105820974944592307816406286")
	},
}
