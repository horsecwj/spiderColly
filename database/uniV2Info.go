package database

import (
	"fmt"
	"strings"
)

type UniV2Info struct {
	PairIndex int64  // 交易所在的 PairIndex
	Address   string // 区块信息
	Token0    string
	Token1    string
	Success   bool   // 解析成功状态
	AddrValue string `json:"addr_value" gorm:"type:longtext"`
	BlockNum  int64
}

//type BybitArticle struct {
//	//gorm.Model
//	Link        string `json:"link" gorm:"unique_index;"`
//	Id          uint32 `json:"id" gorm:"autoincrement;"`
//	Title       string `json:"title"`
//	OverView    string `json:"over_view"`
//	Article     string `json:"article" gorm:"type:longtext"`
//	Articletext string `json:"articletext" gorm:"type:longtext"`
//
//	Time      string `json:"time"`
//	Timestamp int64  `json:"timestamp"`
//}
func (db *DBConn) SaveUniV2(array []UniV2Info) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?,?,?)")

		params = append(params, address.PairIndex, address.Address)
		params = append(params, address.Token0, address.Token1)
		params = append(params, address.AddrValue, address.BlockNum)
	}

	// 拼接SQL
	format := "insert into uni_v2_info (pair_index,address,token0,token1,addr_value,block_num) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// MaxBlockNumber 查询当前已获取的最大区块高度
func (db *DBConn) MaxUniV2Number() (uint, error) {

	sql := "select ifnull(max(pair_index), 0) pair_index from uni_v2_info"
	var numbers []uint

	err := db.Raw(sql).Pluck("pair_index", &numbers).Error
	if err != nil {
		return 0, err
	}

	return numbers[0], nil
}
