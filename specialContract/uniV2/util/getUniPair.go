package zrxErc20

import (
	store "Spider/specialContract/uniV2/uniPair"
	"github.com/ethereum/go-ethereum/common"
	"log"
)

func UseUniPairDemo(address common.Address) (common.Address, common.Address, int64) {
	client := EVMInstance().GetClinet()
	instance, err := store.NewUniPair(address, client)
	if err != nil {
		log.Fatal(err)
	}
	token1, err := instance.Token0(nil)
	if err != nil {
		log.Fatal(err)
	}
	token2, err := instance.Token1(nil)
	bal, err := instance.BalanceOf(nil, address)
	if err != nil {
		return [20]byte{}, [20]byte{}, 0
	}
	return token1, token2, bal.Int64()
}
