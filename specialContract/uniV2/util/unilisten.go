package zrxErc20

import (
	"Spider/config"
	exchange "Spider/specialContract/uniV2/uniPair" // for demo
	ethUtil "Spider/spiderService/ethereum/util"
	"context"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/spf13/cast"
	"log"
	"math/big"
	"strings"
)

var (
	EVMConfig *config.ETHConfig
)

func EVMInstance() *ethUtil.ETHClient {
	c := EVMConfig
	ethUtil.EVMInstance(c.Chain)
	return ethUtil.EVMInstance(EVMConfig.Chain)
}

type UniPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

//0xBb2b8038a1640196FbE3e38816F3e67Cba72D940 ca
func GetUniTrans(blockumber int64, ca string, addrMap map[string]bool) (error, bool) {
	var (
		err  error
		over bool
	)
	client := EVMInstance().GetClinet()
	contractAddress := common.HexToAddress(ca)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(blockumber),
		ToBlock:   big.NewInt(blockumber + intervItme),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	//log.Print("blockumber",blockumber)
	//log.Print("blockumber - interTime",blockumber-intervItme)

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("len(logs) :   ", len(logs))
	contractAbi, err := abi.JSON(strings.NewReader(exchange.UniPairABI))
	if err != nil {
		log.Fatal(err)
	}
	logFillEvent := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	var allTrans []UniPairTransfer
	for _, vLog := range logs {
		if over {
			break
		}
		item := vLog.TxHash
		log.Print(item)
		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			var fillEvent UniPairTransfer
			res1 := make(map[string]interface{})
			err := contractAbi.Events["Transfer"].Inputs.UnpackIntoMap(res1, vLog.Data)
			if err != nil {
				log.Fatal(err)
				return err, over
			}
			fillEvent.From = common.HexToAddress(vLog.Topics[1].String())
			fillEvent.To = common.HexToAddress(vLog.Topics[2].String())
			addrMap[vLog.Topics[1].String()] = true
			addrMap[vLog.Topics[2].String()] = true
			fillEvent.Value = res1["value"].(*big.Int)
			allTrans = append(allTrans, fillEvent)
		}
		rec, err := client.TransactionReceipt(context.Background(), vLog.TxHash)
		if err != nil {
			return err, over
		}
		rca := rec.ContractAddress.String()
		rca = strings.ToLower(rca)
		ca = strings.ToLower(ca)
		if rca == ca {
			over = true
		}

	}
	return err, over
}

func GetBtweenLog(ca string, addrMap map[string]bool) error {

	_, err := EVMInstance().LastBlockNumber()
	if err != nil {
		return err
	}
	//log.Print(blockNum)
	for i := 13897211; i > 0; i -= intervItme {
		log.Print("i", i)
		err, over := GetUniTrans(cast.ToInt64(i-intervItme), ca, addrMap)
		if err != nil {
			return err
		}
		if over {
			break
		}
	}
	return err
}
