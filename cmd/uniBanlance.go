package cmd

import (
	"Spider/common"
	"Spider/config"
	"Spider/database"
	se "Spider/specialContract"
	ethUtil "Spider/spiderService/ethereum/util"
	"github.com/spf13/cobra"
	"os"
)

var uniBalanceCommand = &cobra.Command{
	Use:   "uniBalance",
	Short: "run ethereum service",
	Run: func(cmd *cobra.Command, args []string) {
		err := database.Init(false)
		// 关闭连接
		defer database.CloseConn()
		// 初始化Logger
		common.InitLogger("uniBalance")
		database.AutoMigrate()
		if err != nil {
			common.Logger.Info("uniBalance -> ", err)
			os.Exit(1)
			return
		}
		ethUtil.InitETHInstance()
		if err := se.Run(config.ETHConf()); err != nil {
			common.Logger.Info("uniBalance service 启动失败:%s", err)
			os.Exit(1)
			return
		}
	},
}

func init() {

	rootCommand.AddCommand(uniBalanceCommand)
}
