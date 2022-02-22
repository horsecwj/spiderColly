package schedule

import (
	"Spider/common"
	"Spider/database"
	"Spider/spiderService/util"
	"github.com/robfig/cron"
	"log"
	"time"
)

func createTopGameFi(c *cron.Cron) (err error) {

	err = c.AddFunc("@every 10m", synCmcGameFi)
	if err != nil {
		log.Print(err)
		return
	}
	err = c.AddFunc("@every 10m", syncCoinGk)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("createTopGameFi success ")

	return
}

func synCmcGameFi() {

	err := util.Retry(3, 1*time.Second, cmcGameFi)
	if err != nil {
		common.Logger.Info("更新失败 synCmcGameFi:", err)
		log.Print(err)
	}
	log.Print("synCmcGameFi")
}

func syncCoinGk() {

	err := util.Retry(3, 1*time.Second, cgkGameFi)
	if err != nil {
		common.Logger.Info("更新失败 syncCoinGk :", err)
		log.Print(err)
	}
	log.Print("syncCoinGk")
}

func cgkGameFi() error {

	db := database.DB()
	err := db.DeleteTopCkoGameFi()
	if err != nil {
		common.Logger.Info("删除TopCkoGameFi出错:", err)
		return err
	}

	res, err := util.GetTopGameFiCoinCko()
	if err != nil {
		common.Logger.Info("获取TopCkoGameFi:", err)
		return err
	}

	err = db.SaveTopCkoGameFi(res)
	if err != nil {
		common.Logger.Info("插入TopCkoGameFi:", err)
		return err
	}
	return nil
}

func cmcGameFi() error {
	db := database.DB()
	err := db.DeleteTopCmkGameFi()
	if err != nil {
		common.Logger.Info("删除TopCmkGameFi出错:", err)
		return err
	}

	res, err := util.GetTopGameFiCoinMarket()
	if err != nil {
		common.Logger.Info("获取TopCmkGameFi:", err)
		return err
	}

	err = db.SaveTopCmkGameFi(res[:10])
	if err != nil {
		common.Logger.Info("插入TopCmkGameFi:", err)
		return err
	}
	return nil
}
