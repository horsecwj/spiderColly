package evmConfig

import "Spider/config"

var (
	EVMConfig *config.ETHConfig
)

func InitEvm(conf *config.ETHConfig) {
	EVMConfig = conf
}
