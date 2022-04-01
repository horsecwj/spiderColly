package util

import (
	"Spider/common"
	"Spider/config"
	"Spider/database"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"os"
	"testing"
)

func initDatabse() {
	viper.AddConfigPath("../../../config")
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
	// 尝试初始化数据库连接
	err := database.Init(true)
	// 关闭连接
	//defer database.CloseConn()

	if err != nil {

		common.Logger.Info("数据库初始化失败:", err)
		os.Exit(1)
		return
	}
}

func TestUseUniV3Fac(t *testing.T) {
	initDatabse()

	UseUniV3Fac()
}
