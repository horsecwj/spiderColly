package database

import (
	"fmt"
	"strings"
)

type UniV3Info struct {
	PairIndex int
	Tk0       string
	Tk1       string
	LikQty    int64
	Fee       int64
	Owner     string
}

func (db *DBConn) SaveUniV3(array []UniV3Info) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?,?, ?, ?, ?,?)")

		params = append(params, address.PairIndex, address.Tk0, address.Tk1)
		params = append(params, address.LikQty, address.Fee)
		params = append(params, address.Owner)
	}

	// 拼接SQL
	format := "insert into uni_v3_info (pair_index,tk0,tk1,lik_qty,fee,owner) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// MaxBlockNumber 查询当前已获取的最大区块高度
func (db *DBConn) MaxUniV3Number() (uint, error) {

	sql := "select ifnull(max(pair_index), 0) pair_index from uni_v3_info"
	var numbers []uint

	err := db.Raw(sql).Pluck("pair_index", &numbers).Error
	if err != nil {
		return 0, err
	}

	return numbers[0], nil
}
