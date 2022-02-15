package cmd

import (
	"Spider/config"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"os"
)

var configPath string

var rootCommand = &cobra.Command{

	Use:   "blockchain-wallet",
	Short: "run blockchain wallet",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func init() {

	rootCommand.PersistentFlags().StringVar(&configPath, "config-path", "./config", "system config  path")
	cobra.OnInitialize(initConfig)

	_ = viper.BindPFlag("config-path", rootCommand.PersistentFlags().Lookup("config-path"))
}

func initConfig() {

	viper.AddConfigPath(configPath)
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("配置读取出错: ", err)
		return
	}

	// 监听配置
	viper.OnConfigChange(func(in fsnotify.Event) {
		config.RefreshConf()
	})
	viper.WatchConfig()
}

func Execute() {

	err := rootCommand.Execute()
	if err != nil {
		fmt.Println("启动失败: ", err)
		os.Exit(1)
	}
	log.Print(111)
}
