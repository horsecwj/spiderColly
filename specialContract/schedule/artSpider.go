package schedule

import (
	sc2 "Spider/specialContract/uniV2/util"
	sc3 "Spider/specialContract/uniV3/util"
	"github.com/robfig/cron"

	"log"
)

func createUniV2(c *cron.Cron) (err error) {

	err = c.AddFunc("@every 240h", sc2.CronUniV2)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("CronUniV2 success ")
	return
}

func createUniV3(c *cron.Cron) (err error) {

	err = c.AddFunc("@every 240h", sc3.UseUniV3Fac)
	if err != nil {
		log.Print(err)
		return
	}
	log.Printf("UseUniV3Fac success ")
	return
}
