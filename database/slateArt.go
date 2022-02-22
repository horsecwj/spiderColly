package database

import (
	"Spider/spiderService/model"
	"fmt"
	"log"
	"strings"
)

// 批量保存
func (db *DBConn) SaveSlateArt(array []model.SlateArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?,?)")
		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time, address.Timestamp)
	}

	format := "insert into slate_article (title,over_view,article,link,time,timestamp) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

// 获取一个未使用的地址
func (db *DBConn) GetSlateArt() ([]*model.SlateArticle, error) {
	var addr []*model.SlateArticle
	err := db.Model(&addr).Debug().Order("timestamp desc limit 1").Scan(&addr).Error
	log.Print(err)
	return addr, err
}
