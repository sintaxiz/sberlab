package cmd

import (
	"fmt"
	"os"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "cloudmanager",
	Short: "",
	Long:  ``,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var key string
var secret string
var cfgFile string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "Config file. Default is current directory")

	rootCmd.PersistentFlags().StringVar(&key, "key", "", "Authentication key for singing")
	viper.BindPFlag("key", rootCmd.PersistentFlags().Lookup("key"))

	rootCmd.PersistentFlags().StringVar(&secret, "secret", "", "Secret key. Needs for authentication.")
	viper.BindPFlag("secret", rootCmd.PersistentFlags().Lookup("secret"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
		fmt.Println("Program will use your config...")
	} else {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigName("cloudmanager.json")
	}
	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Can not read config:", err)
	}
	key = viper.GetString("key")
	secret = viper.GetString("secret")
}
