package schedule

import (
	"Spider/config"
	"Spider/specialContract/evmConfig"
	"github.com/robfig/cron"
)

// 创建定时执行任务
func create() (c *cron.Cron, err error) {

	c = cron.New()

	// 添加article任务执行
	err = createUniV2(c)
	if err != nil {

		return
	}

	// 添加topGameFi任务同步
	err = createUniV3(c)
	if err != nil {

		return
	}

	return
}

func Run(conf *config.ETHConfig) (err error) {

	evmConfig.InitEvm(conf)
	c, err := create()

	if err != nil {

		return
	}

	c.Run()
	return
}
