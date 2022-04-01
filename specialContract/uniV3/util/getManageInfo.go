package util

import (
	sd "Spider/database"
	store "Spider/specialContract/uniV3/NonfungiblePositionManager"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
	"log"
	"math/big"
)

func UseUniV3Fac() {

	address := common.HexToAddress("0xC36442b4a4522E871399CD717aBDD847Ab11FE88")
	pairStartIndex, err := sd.DB().MaxUniV3Number()
	if err != nil {
		return
	}
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := store.NewUniV3(address, client)
	if err != nil {
		log.Fatal(err)
	}
	length, err := instance.TotalSupply(nil)
	if err != nil {
		log.Fatal(err)
	}
	lengthNum := cast.ToInt(length.Int64())
	for itemIndex := int(pairStartIndex) + 1; itemIndex <= lengthNum-1; itemIndex++ {

		pos, err := instance.Positions(nil, big.NewInt(cast.ToInt64(itemIndex)))
		if err != nil {
			continue
		}
		oweer, err := instance.OwnerOf(nil, big.NewInt(cast.ToInt64(itemIndex)))
		if err != nil {
			continue
		}
		log.Print(oweer, pos)
		temp := sd.UniV3Info{
			PairIndex: itemIndex,
			Tk0:       pos.Token0.String(),
			Tk1:       pos.Token1.String(),
			LikQty:    pos.Liquidity.Int64(),
			Fee:       pos.Fee.Int64(),
			Owner:     oweer.String(),
		}
		var ts []sd.UniV3Info
		ts = append(ts, temp)
		_ = sd.DB().SaveUniV3(ts)
	}
}
