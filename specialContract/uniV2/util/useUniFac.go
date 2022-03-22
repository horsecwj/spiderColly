package zrxErc20

import (
	store "Spider/specialContract/uniV2/contractFactory"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

func UseUniFac() {

	address := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")

	client := EVMInstance().GetClinet()
	instance, err := store.NewUniFac(address, client)

	if err != nil {
		log.Fatal(err)
	}
	length, err := instance.AllPairsLength(nil)
	if err != nil {
		log.Fatal(err)
	}

	pair, err := instance.AllPairs(nil, big.NewInt(1000))
	addr := pair.Hex()
	log.Print(length, addr)

}
