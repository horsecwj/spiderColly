package database

import (
	"Spider/spiderService/model"
	"fmt"
	"strings"
)

func (db *DBConn) SaveBybitNewlyArt(array []model.BybitNewlyArticle) error {

	if len(array) == 0 {

		return nil
	}

	values := make([]string, 0, len(array))
	params := make([]interface{}, 0, len(array)*7)
	for _, address := range array {

		values = append(values, "(?, ?, ?, ?, ?)")
		params = append(params, address.Title, address.OverView)
		params = append(params, address.Article, address.Link)
		params = append(params, address.Time)
	}

	format := "insert into bybit_newly_article (title,over_view,article,link,time) values %s"
	sql := fmt.Sprintf(format, strings.Join(values, ","))

	return db.Exec(sql, params...).Error
}

func (db *DBConn) GetBybitNewlyArt() ([]*model.BybitNewlyArticle, error) {
	var addr []*model.BybitNewlyArticle
	err := db.Model(&addr).Debug().Order("id desc limit 1").Scan(&addr).Error
	return addr, err
}

func (db *DBConn) GetBybitNewlyArtBy() (model.BybitNewlyArticle, error) {
	var addr model.BybitNewlyArticle
	err := db.Model(&addr).Debug().Where(&model.BybitNewlyArticle{Id: 24}).Scan(&addr).Error
	return addr, err
}
