package schedule

import (
	"Spider/common"
	"Spider/database"
	"Spider/spiderService/model"
	"Spider/spiderService/util"
	"github.com/robfig/cron"
	"log"
	"sort"
	"time"
)

func createArtSpider(c *cron.Cron) (err error) {

	err = c.AddFunc("@every 2h", syncBybit)
	if err != nil {
		log.Print(err)
		return
	}
	err = c.AddFunc("@every 2h", synCmc)
	if err != nil {
		log.Print(err)
		return
	}
	return
}

func synCmc() {
	log.Print("syncBycmc")
	err := util.Retry(3, 1*time.Second, cmcArt)
	if err != nil {
		common.Logger.Info("更新失败cmcart:", err)
		log.Print(err)
	}
}

func syncBybit() {
	log.Print("syncBybit")
	err := util.Retry(3, 1*time.Second, bybitHighly)
	if err != nil {
		common.Logger.Info("更新失败bybitart:", err)
		log.Print(err)
	}
	err = util.Retry(3, 1*time.Second, bybitNewly)

	if err != nil {
		common.Logger.Info("更新失败bybitnewlyart:", err)

		log.Print(err)
	}

}

func bybitHighly() error {
	var (
		err  error
		resM []model.BybitArticle
	)
	db := database.DB()
	res, err := db.GetBybitArt()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		resM, err = util.GetArticleBybitArt("res.Link")

	} else {
		resM, err = util.GetArticleBybitArt(res[0].Link)

	}

	if err != nil {
		return err
	}
	sort.Sort(util.BybitArticleSlice(resM))
	err = db.SaveBybitHighLightArt(resM)
	if err != nil {
		return err
	}
	return nil
}

func bybitNewly() error {
	var (
		err  error
		resM []model.BybitNewlyArticle
	)
	db := database.DB()
	res, err := db.GetBybitNewlyArt()
	if err != nil {
		return err
	}
	if len(res) == 0 {
		resM, err = util.GetNewArticleBybitArt("res[0].Link")
	} else {
		resM, err = util.GetNewArticleBybitArt(res[0].Link)
	}
	if err != nil {
		return err
	}
	sort.Sort(util.BybitNewlyArticleSlice(resM))
	err = db.SaveBybitNewlyArt(resM)
	if err != nil {
		return err
	}
	return nil
}

func cmcArt() error {
	var (
		err  error
		resM []model.SlateArticle
	)
	db := database.DB()
	res, err := db.GetSlateArt()
	if err != nil {
		common.Logger.Info("database 获取失败:", err)
		return err
	}

	if len(res) != 0 {
		resM, err = util.GetArticleCryptoSlate(res[0].Link)
	} else {
		resM, err = util.GetArticleCryptoSlate("res[0].Link")
	}
	if err != nil {
		common.Logger.Info(" 爬取失败:", err)
		return err
	}
	sort.Sort(util.SlateArticleSlice(resM))
	err = db.SaveSlateArt(resM)
	if err != nil {
		common.Logger.Info(" 插入失败:", err)
		return err
	}
	return nil
}
