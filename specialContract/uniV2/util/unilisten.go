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

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(exchange.UniPairABI))
	if err != nil {
		log.Fatal(err)
	}
	//addrMap = make(map[string]bool)
	logFillEvent := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	for _, vLog := range logs {
		if over {
			break
		}
		var fillEvent UniPairTransfer
		switch vLog.Topics[0].Hex() {

		case logFillEvent.Hex():

			res1 := make(map[string]interface{})
			err := contractAbi.Events["Transfer"].Inputs.UnpackIntoMap(res1, vLog.Data)
			if err != nil {
				log.Fatal(err)
				return err, over
			}
			addr1 := vLog.Topics[1].String()
			addr2 := vLog.Topics[2].String()

			fillEvent.From = common.HexToAddress(addr1)
			fillEvent.To = common.HexToAddress(addr2)
			addrMap[addr1] = true
			addrMap[addr2] = true
			fillEvent.Value = res1["value"].(*big.Int)

			tx := vLog.TxHash
			rec, err := client.TransactionReceipt(context.Background(), tx)
			if err != nil {
				return err, over
			}
			rca := rec.ContractAddress.String()
			rca = strings.ToLower(rca)
			ca = strings.ToLower(ca)
			if rca == ca {
				over = true
			}
			//log.Print("tx : ",tx)
		}
	}
	return err, over
}

func GetBtweenLog(worker *TransactionWorker) error {

	ca := worker.Address

	//addrMap := make(map[string]bool)
	//worker.AddrMap
	blockNum, err := EVMInstance().LastBlockNumber()
	if err != nil {
		return err
	}
	worker.BlockNum = cast.ToInt64(blockNum)
	//log.Print(blockNum) 10007197
	for i := blockNum; i > 10007197; i -= intervItme {
		log.Printf("index : %d, i :%d ", worker.Index, i)
		if i <= intervItme {
			i = intervItme
		}
		err, over := GetUniTrans(cast.ToInt64(i-intervItme), ca, worker.AddrMap)

		if err != nil {
			return err
		}
		if over {
			break
		}
	}

	log.Print("over download ", worker.Address)
	return err
}
