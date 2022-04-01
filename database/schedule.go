package database

import (
	"Spider/common"
	"Spider/common/types"
	"Spider/config"
	ethUtil "Spider/spiderService/ethereum/util"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/spf13/cast"

	//"Spider/config"
	"encoding/hex"
	"fmt"
	ecommon "github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/common/hexutil"
	ethTypes "github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"strings"
	//"Spider/database"
	"sync"
)

func EVMInstance() *ethUtil.ETHClient {

	return ethUtil.EVMInstance(config.ETHConf().Chain)
}

type GetBlockNumber func() (uint, error)

type SyncBlockScheduler struct {
	Height           uint           // 区块确认高度
	ChainType        types.Chain    // 公链类型
	ChainBlockNumber GetBlockNumber // 获取区块链的当前高度
}

//
func (scheduler *SyncBlockScheduler) Start() {

	// 获取当前区块高度
	number, err := scheduler.ChainBlockNumber()

	if err != nil {

		common.Logger.Info("查询当前区块高度失败:", err)
		return
	}

	// 计算可同步高度
	current := number - scheduler.Height
	LoclaUnhandle, err := DB().MaxUnHandleBlockNumber(scheduler.ChainType)
	if err != nil {

		common.Logger.Info("查询本地区块高度失败:", err)
		return
	}
	CountUnHandleBlockNumber, err := DB().CountUnHandleBlockNumber(scheduler.ChainType)
	if err != nil {

		common.Logger.Info("查询本地区块高度失败:", err)
		return
	}
	if LoclaUnhandle == 0 {
		LoclaUnhandle = 100000
	}

	if CountUnHandleBlockNumber >= 100 {

		return
	}

	if current >= LoclaUnhandle+500 {

		current = LoclaUnhandle + 500
	}

	common.Logger.Info("发现新高度:", current, " 当前高度:", LoclaUnhandle)

	// 保存至数据库
	err = DB().SaveBlockNumbers(LoclaUnhandle, current, scheduler.ChainType)
	if err != nil {

		common.Logger.Info("保存区块解析任务失败: ", err)
	}
	return
}

type SyncTransactionScheduler struct {
	WaitGroup    *sync.WaitGroup
	BlockLimit   uint
	ChainType    types.Chain
	BlockHandler BlockInfoHandler
	TxHandler    TransactionHandler
}

type TransactionWorker struct {
	Wait      *sync.WaitGroup
	Index     int            // 交易所在的 Index
	BlockInfo *BlockInfo     // 区块信息
	Recharges []*Transaction // 解析出来的充值记录
	Success   bool           // 解析成功状态
	CkInfo    *TransInfo
}

type BlockInfo struct {
	Number     int
	Hash       string
	Timestamp  int64
	TxSize     int
	Getter     TransactionGetter
	Difficulty int64
}

type MyBlockInfo struct {
	Number     int
	Hash       string
	Timestamp  int64
	TxSize     int
	Difficulty int64
}

type TransactionGetter func(index int) (tx interface{})

type BlockInfoHandler func(blockNumber uint) (info *BlockInfo, err error)

type TransactionHandler func(worker *TransactionWorker)

func (scheduler *SyncTransactionScheduler) Start() {

	array, err := DB().GetUnhandledBlock(scheduler.ChainType, scheduler.BlockLimit)
	if err != nil {

		common.Logger.Info("查询未解析区块任务失败:", err)
		return
	}

	if len(array) == 0 {

		return
	}

	scheduler.WaitGroup.Add(len(array))
	for _, block := range array {

		go scheduler.parseBlock(*block)

	}
}

//
//// 解析区块
func (scheduler *SyncTransactionScheduler) parseBlock(item BcwBlockNumber) {

	defer func() {
		scheduler.WaitGroup.Done()
	}()

	info, err := scheduler.BlockHandler(uint(item.Number))
	if err != nil {
		return
	}

	blocksInfo := MyBlockInfo{Number: info.Number, Hash: info.Hash, Timestamp: info.Timestamp, TxSize: info.TxSize, Difficulty: info.Difficulty}
	blocksInfo.InsertTransTable("block")

	// 分页解析（单次解析数量过多容易导致数据库连接耗尽)
	limit := 100
	page := info.TxSize / limit
	if info.TxSize%limit > 0 {

		page += 1
	}

	pageLimit := limit
	if info.TxSize < limit {
		pageLimit = info.TxSize
	}

	tasks := make([]*TransactionWorker, 0, info.TxSize)
	tasksInfo := make([]*TransInfo, 0, info.TxSize)
	for i := 0; i < page; i++ {

		var txWait sync.WaitGroup

		if i == page-1 && info.TxSize%pageLimit > 0 {

			txWait.Add(info.TxSize % pageLimit)
		} else {

			txWait.Add(pageLimit)
		}

		// DO Something
		for j := 0; j < pageLimit; j++ {

			index := i*pageLimit + j
			if index >= info.TxSize {

				break
			}
			worker := &TransactionWorker{
				Wait:      &txWait,
				Index:     index,
				BlockInfo: info,
			}
			go func(workertemp *TransactionWorker) {
				temp := ParseTransaction(workertemp)
				if len(temp.TId) != 0 || temp.Value != 0 {
					temp.InsertTransTable("transaction")
					tasksInfo = append(tasksInfo, &temp)
				}
			}(worker)
			tasks = append(tasks, worker)
		}
		txWait.Wait()
	}

	//err = db.SaveEthTrans(tasksInfo)
	//if err != nil{
	//	log.Print(err)
	//}
	// 检查解析状态
	recharges, succeed := scheduler.checkParseStatus(tasks)

	// 全部解析成功
	if succeed {
		// 保存充值交易
		err = scheduler.blockParseSucceed(&item, recharges)
		if err != nil {
			common.Logger.Info("更新解析状态失败:", err)
		}
	}

}

func ParseTransaction(worker *TransactionWorker) TransInfo {
	defer func() {
		worker.Success = true
		worker.Wait.Done()
	}()

	tmp := TransInfo{
		Hash:             "",
		BlockHash:        "",
		Nonce:            0,
		BlockNumber:      0,
		TransactionIndex: 0,
		FromAddr:         "",
		ToAddr:           "",
		Value:            0,
		Gas:              0,
		GasPrice:         0,
		BlockTimestamp:   0,
		Data:             "",
	}

	tx := worker.BlockInfo.Getter(worker.Index).(*ethTypes.Transaction)

	//client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	//txHash := ecommon.HexToHash("0x806d1e8d6bde0539cbbc4a966906642bd84830e33f68b931bfd3938b135a7c16")
	//worker.BlockInfo.Hash = "0xa1d74c520b5a71c02acfa3a62947be49d88e496d6abd07cf2cf2033045fbd8bf"
	////worker.Index =
	//tx, _, err  = client.TransactionByHash(context.Background(), txHash)
	//if err != nil {
	//	log.Fatal(err)
	//}

	receiver := tx.To()
	if receiver == nil {
		worker.Success = true
		common.Logger.Info(tx.Hash().Hex(), "：无接收地址")
		return TransInfo{}
	}
	receiverAddr := receiver.Hex()

	q := tx.Value()
	quantity := q.Uint64()
	data := tx.Data()
	var tid string
	if data != nil && len(data) >= 4+32*2 {
		rec, tidInt, qua, _, err := parseInputData(tx.Hash().String(), receiverAddr, data)
		tid = cast.ToString(tidInt)
		if err == nil && !strings.EqualFold(rec, "") {
			//symbol, _ = DB().GetSymbol(sym)
			receiverAddr = rec
			quantity = qua
		} else {
			worker.Success = true
		}
	} else {
		worker.Success = true
	}

	sender, err := EVMInstance().GetSender(tx, worker.BlockInfo.Hash, uint64(worker.Index))
	if err != nil {
		common.Logger.Warn("获取交易发送者失败:", err)
		return tmp
	}

	if sender.Hex() == "0x0000000000000000000000000000000000000000" {

		common.Logger.Error("获取sender 错误 -> 0x0000000000000000000000000000000000000000")
		return tmp
	}

	temp := TransInfo{
		TId:              tid,
		Hash:             tx.Hash().Hex(),
		BlockHash:        worker.BlockInfo.Hash,
		Nonce:            int64(tx.Nonce()),
		BlockNumber:      worker.BlockInfo.Number,
		TransactionIndex: worker.Index,
		FromAddr:         sender.String(),
		ToAddr:           receiverAddr,
		Value:            int64(quantity),
		Data:             string(tx.Data()),
		Gas:              int64(tx.Gas()),
		GasPrice:         tx.GasPrice().Int64(),
		BlockTimestamp:   worker.BlockInfo.Timestamp,
	}
	worker.CkInfo = &temp
	worker.Success = true
	return temp

}

func parseInputData(txHash, contract string, data []byte) (receiver string, tid, quantity uint64, symbol string, err error) {

	{
		mthod := data[:4]
		ms := hexutil.Encode(mthod)
		var address ecommon.Address
		//var tid uint64
		switch ms {
		///erc20 transfer
		case "0xa9059cbb": //0xec35097177d51ebe983dcb72b3d2d01b88cb5394cf7ca767868801898c511a7f
			balanceMethod := "70a08231"
			// 再截取地址
			addressData := data[4:36]
			address = ecommon.BytesToAddress(addressData[12:])
			valid := validateERC20Balance(contract, balanceMethod, address.Hex(), big.NewInt(0))
			if !valid {
				err = fmt.Errorf("ERC20 代币金额错误")
				return
			}
			quantity = ParseDataToUint(data[36:68])

		case "0x42842e0e": //0x8f92d5bfe185d3f1f888e52f4d6cd225fda00066a1c645e9b74064f723cfa9fc  safe Transefer
			addressData := data[36:68]
			address = ecommon.BytesToAddress(addressData[12:])
			tid = ParseDataToUint(data[68:100])
			quantity = 1
		case "0x23b872dd": //	0x806d1e8d6bde0539cbbc4a966906642bd84830e33f68b931bfd3938b135a7c16 transefer from
			addressData := data[36:68]
			address = ecommon.BytesToAddress(addressData[12:])
			tid = ParseDataToUint(data[68:100])
			quantity = 1
		case "0xf242432a": //0x4766a1731beaa5640a66e73524e772dd4ff297de777358598361c110c02b4887 safeTransferFrom
			addressData := data[36:68]
			address = ecommon.BytesToAddress(addressData[12:])
			tid = ParseDataToUint(data[68:100])
			quantity = ParseDataToUint(data[100:132])
		}

		// 判断执行结果
		receipt, e := EVMInstance().GetTransactionReceipt(txHash)
		if e != nil {
			err = e
			common.Logger.Warn("查询receipt出错:", txHash, "----:", err)
			return
		}
		if receipt.Status != 1 {
			common.Logger.Warn("合约执行失败不处理:", txHash)
			return
		}
		receiver = address.Hex()
		//logs := receipt.Logs
		//for _,vLog := range logs{
		//
		//	switch vLog.Topics[0].Hex() {
		//	case "0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"{
		//
		//	}
		//
		//	}
		//
		//}

		return
	}
}

func ParseDataToUint(idData []byte) uint64 {
	//idData := data[68:100]
	var err error
	id := hex.EncodeToString(idData)
	// 去除前缀的0，并添加0x
	id = parseInputDataQuantityHex(id)
	var b *big.Int
	b, err = hexutil.DecodeBig(id)
	if err != nil {
		//common.Logger.Warn(txHash, " - 解析 ERC20 金额失败:", err)
		return 0
	}
	tId := b.Uint64()
	return tId
}

func parseInputDataQuantityHex(value string) string {

	if strings.HasPrefix(value, "0") {

		return parseInputDataQuantityHex(strings.Replace(value, "0", "", 1))
	} else {

		return "0x" + value
	}
}

// 验证
func validateERC20Balance(contract string, method string, address string, amount *big.Int) bool {
	balance, err := EVMInstance().GetERC20Balance(contract, method, address)
	if err != nil {

		common.Logger.Warn("解析ERC20余额出错:", err)
		return false
	}

	return balance.Cmp(amount) >= 0
}

// 解析成功更新状态
func (scheduler *SyncTransactionScheduler) blockParseSucceed(block *BcwBlockNumber, recharges []*Transaction) error {

	db := DB()

	// 没有本平台的充值记录
	if len(recharges) == 0 {

		return block.ParseComplete()
	}

	_, err := db.Transaction(func(tx *DBConn) (i interface{}, e error) {

		//e = tx.SaveTransactions(recharges, scheduler.ChainType)
		//if e != nil {
		//	return
		//}

		e = block.ParseCompleteInDB(tx)

		return
	})

	return err
}

// 检查解析结果
func (scheduler *SyncTransactionScheduler) checkParseStatus(tasks []*TransactionWorker) (array []*Transaction, succeed bool) {

	succeed = true
	//
	for _, worker := range tasks {

		// 出现解析失败
		if !worker.Success {

			succeed = false
			break
		}

		if worker.Recharges != nil && len(worker.Recharges) > 0 {

			array = append(array, worker.Recharges...)
		}
	}

	return
}

type SyncRechargeScheduler struct {
	ChainType        types.Chain    // 公链类型
	ChainBlockNumber GetBlockNumber // 获取区块链的当前高度
	Confirmations    uint64         // 不可逆高度差
}

type SyncFailedTxScheduler struct {
	Wait      *sync.WaitGroup
	ChainType types.Chain
	Handler   FailedTxHandler
	Limit     uint
}

type FailedTxHandler func(wait *sync.WaitGroup, transaction *FailedTransaction)

func (scheduler *SyncFailedTxScheduler) Start() {

	array, err := DB().FailedTxWithType(scheduler.ChainType, scheduler.Limit)
	if err != nil {

		common.Logger.Info("查询", scheduler.ChainType, "失败交易错误:", err)
		return
	}

	// 没有失败交易
	if array == nil || len(array) == 0 {

		return
	}

	scheduler.Wait.Add(len(array))

	for _, tx := range array {

		go scheduler.Handler(scheduler.Wait, tx)
	}
}

type SyncStatusScheduler struct {
	ChainType        types.Chain
	Confirmations    uint64
	CheckInterval    int64
	ChainBlockNumber GetBlockNumber
	TxStatus         StatusHandler
}

type StatusHandler func(txHash string) (exists bool, err error)
