package ethereum

import (
	"Spider/config"
	"Spider/spiderService/ethereum/schedule"
)

func Run(conf *config.ETHConfig) error {

	return schedule.Run(conf)
}
