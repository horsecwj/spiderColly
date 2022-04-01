package zrxErc20

import (
	//sc "Spider/common"
	//"Spider/database"
	sd "Spider/database"
	store "Spider/specialContract/uniV2/contractFactory"
	storePair "Spider/specialContract/uniV2/uniPair"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cast"
	"log"
	"math/big"
	"sync"
)

type TransactionWorker struct {
	Wait      *sync.WaitGroup
	Index     int64 // 交易对 所在的 所有的交易对 的Index
	Address   string
	Token0    string
	Token1    string
	Success   bool // 解析成功状态
	AddrMap   map[string]bool
	AddrValue map[string]int64
	BlockNum  int64
}

func GetContribution(address common.Address) {

	startIndex, err := sd.DB().MaxUniV2Number()
	if err != nil {
		return
	}

	client := EVMInstance().GetClinet()
	instance, err := store.NewUniFac(address, client)
	if err != nil {
		log.Fatal(err)
	}
	//所有的交易对
	length, err := instance.AllPairsLength(nil)
	if err != nil {
		log.Fatal(err)
	}
	lengthNum := cast.ToInt(length.Int64())
	limit := 1
	page := lengthNum / limit
	if lengthNum > 0 {
		page += 1
	}

	pageLimit := limit
	if lengthNum < limit {
		pageLimit = lengthNum
	}

	tasks := make([]*TransactionWorker, 0, lengthNum)

	for i := 0; i < page; i++ {

		var txWait sync.WaitGroup

		if i == page-1 && lengthNum%pageLimit > 0 {

			txWait.Add(lengthNum % pageLimit)
		} else {

			txWait.Add(pageLimit)
		}

		// DO Something
		for j := 0; j < pageLimit; j++ {

			index := i*pageLimit + int(startIndex) + 1
			if index >= lengthNum {

				break
			}

			pIndex := index
			pair, err := instance.AllPairs(nil, big.NewInt(cast.ToInt64(pIndex)))
			if err != nil {
				log.Print(err)
				return
			}

			log.Print("pair.String : ", pair.String())

			worker := &TransactionWorker{
				Wait:      &txWait,
				Index:     cast.ToInt64(index),
				Address:   pair.String(),
				AddrMap:   make(map[string]bool),
				AddrValue: make(map[string]int64),
			}

			go func(worker *TransactionWorker) {
				//所有交互地址
				err := GetBtweenLog(worker)
				if err != nil {
					return
				}
				instancePair, err := storePair.NewUniPair(common.HexToAddress(worker.Address), client)
				if err != nil {
					log.Print(err)
				}
				for item := range worker.AddrMap {
					bal, err := instancePair.BalanceOf(nil, common.HexToAddress(item))
					if err != nil {
						log.Print(err)
						continue
					}
					log.Print(bal.Int64())
					worker.AddrValue[item] = bal.Int64()
				}

				t0, err := instancePair.Token0(nil)
				if err != nil {
					log.Print(err)
				} else {
					worker.Token0 = t0.String()
				}
				t1, err := instancePair.Token1(nil)
				if err != nil {
					log.Print(err)
				} else {
					worker.Token1 = t1.String()
				}
				worker.Wait.Done()

			}(worker)
			tasks = append(tasks, worker)
		}
		txWait.Wait()
		lengthIndex := len(tasks) - 1
		dataType, _ := json.Marshal(tasks[lengthIndex].AddrValue)
		dataString := string(dataType)
		temp := sd.UniV2Info{
			PairIndex: tasks[lengthIndex].Index,
			Address:   tasks[lengthIndex].Address,
			Token0:    tasks[lengthIndex].Token0,
			Token1:    tasks[lengthIndex].Token1,
			Success:   false,
			AddrValue: dataString,
			BlockNum:  tasks[lengthIndex].BlockNum,
		}
		var ts []sd.UniV2Info
		ts = append(ts, temp)
		_ = sd.DB().SaveUniV2(ts)
		log.Print("page limit sucess danalyse")
	}
}
