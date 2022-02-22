package database

import (
	"Spider/spiderService/model"
	"fmt"
	"strings"
)

// 批量保存
func (db *DBConn) SaveBybitHighLightArt(array []model.BybitArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)

	// 组装参数
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?,?)")

		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time, address.Timestamp)
	}

	// 拼接SQL
	format := "insert into bybit_article (title,over_view,article,link,time,timestamp) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 获取一个
func (db *DBConn) GetBybitArt() ([]model.BybitArticle, error) {
	var addr []model.BybitArticle
	err := db.Model(&addr).Debug().Order("timestamp  limit 1").Scan(&addr).Error
	return addr, err
}
