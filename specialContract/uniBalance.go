package ethereum

import (
	"Spider/config"
	"Spider/specialContract/schedule"
)

func Run(conf *config.ETHConfig) error {

	return schedule.Run(conf)
}
