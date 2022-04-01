package zrxErc20

import (
	sc "Spider/common"
	"Spider/config"
	"Spider/database"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"testing"
	"time"

	"os"
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

		sc.Logger.Info("数据库初始化失败:", err)
		os.Exit(1)
		return
	}
}

func TestEVMInstance(t *testing.T) {
	//initDatabse()
	////c := config.APPConf()
	//c := config.ETHConf()
	//log.Print(c)
	//EVMConfig = c
	//address := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	//GetContribution(address)
	mainGetUniV2()
}
func mainGetUniV2() {
	defer func() { //catch or finally
		if err := recover(); err != nil { //catch
			log.Printf("Exception: %v\n", err)
			time.Sleep(240 * time.Second)
			mainGetUniV2()
		}
	}()

	initDatabse()
	//c := config.APPConf()
	c := config.ETHConf()
	log.Print(c)
	EVMConfig = c
	address := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	GetContribution(address)

}
