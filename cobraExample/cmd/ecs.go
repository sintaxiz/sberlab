package cmd

import (
	"fmt"

	"example.com/cobraExample/sbercloudAPI"
	"github.com/spf13/cobra"
)

type Ecs struct {
}

var ecsCmd = &cobra.Command{
	Use:   "ecs",
	Short: "Access to elastic cloud servers",
	Long:  `Allow to manage Elastic Cloud Servers`,
	Run: func(cmd *cobra.Command, args []string) {
		sbercloudAPI.SetAuthData(key, secret)
		escJsonArray := sbercloudAPI.GetSubnets()
		fmt.Println(escJsonArray)
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "stop ecs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sbercloudAPI.SetAuthData(key, secret)
		ecsId := "35ba116d-a691-4a0e-a29e-b7b3d810e8a7"
		sbercloudAPI.StopEcs(ecsId)
	},
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start ecs",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		sbercloudAPI.SetAuthData(key, secret)
		ecsId := "35ba116d-a691-4a0e-a29e-b7b3d810e8a7"
		sbercloudAPI.StopEcs(ecsId)
	},
}

func init() {
	ecsCmd.Flags().Bool("subnets", true, "Use Viper for configuration")

	ecsCmd.AddCommand(stopCmd)

	rootCmd.AddCommand(ecsCmd)
}
