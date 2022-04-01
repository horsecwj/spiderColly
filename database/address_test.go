package database

import (
	"Spider/config"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"testing"
)

func TestGetUnusedAddress(t *testing.T) {

	viper.AddConfigPath("../config")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {

		fmt.Println("配置读取出错:", err)
		return
	}
}

func initDatabse() {
	viper.AddConfigPath("../config")
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

}
func TestNewAddress(t *testing.T) {

	initDatabse()
	configt := config.APPConf()
	log.Print(configt)
	CreatTransTable("transaction")
	temp := TransInfo{
		Hash:             "ss",
		BlockHash:        "ss",
		Nonce:            0,
		BlockNumber:      12,
		TransactionIndex: 1,
		FromAddr:         "",
		ToAddr:           "",
		Value:            1,
		Gas:              0,
		GasPrice:         0,
		BlockTimestamp:   0,
		Data:             "",
	}
	temp.InsertTransTable("transaction")
	select {}
}
func TestAutoMigrate(t *testing.T) {

	initDatabse()
	configt := config.APPConf()
	log.Print(configt)
	CreatBlockTable("block")
	temp := MyBlockInfo{
		Number:     0,
		Hash:       "",
		Timestamp:  0,
		TxSize:     0,
		Difficulty: 0,
	}
	temp.InsertTransTable("block")
	select {}
}
