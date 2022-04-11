package zrxErc20

import (
	"github.com/ethereum/go-ethereum/common"
)

func CronUniV2() {
	address := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")
	GetContribution(address)
}
