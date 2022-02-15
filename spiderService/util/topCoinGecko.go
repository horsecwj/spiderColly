package util

import (
	"Spider/spiderService/model"
	"encoding/base64"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func GetTopGameFiCoinCko() ([]*model.TopCkoGameFi, error) {
	var ArrTopGameFi = make([]*model.TopCkoGameFi, 0, 30)
	var err error
	c := colly.NewCollector(
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:96.0) Gecko/20100101 Firefox/96.0"),
	)
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("table tbody tr", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, s *goquery.Selection) {

			if len(ArrTopGameFi) <= 10 {
				str := s.Find("td")
				link, _ := str.Eq(10).Find("img").Attr("data-src")
				res, _ := http.Get(link)
				data, _ := ioutil.ReadAll(res.Body)

				base64Data := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(data)
				log.Println(base64Data)

				tplData := model.TopCkoGameFi{
					ID:        len(ArrTopGameFi),
					Coin:      strings.ReplaceAll(str.Eq(2).Find("a").Eq(0).Text(), "\n", ""),
					Price:     str.Eq(4).Find("span").Text(),
					OneDay:    str.Eq(6).Find("span").Text(),
					OneWeek:   str.Eq(7).Find("span").Text(),
					DayVolume: str.Eq(8).Find("span").Text(),
					MktCap:    str.Eq(9).Find("span").Text(),
					LastWeek:  string(data), //svg+xml文件
				}
				ArrTopGameFi = append(ArrTopGameFi, &tplData)
			}
		})
	})

	err = c.Visit("https://www.coingecko.com/en/categories/gaming")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	c.Wait()
	return ArrTopGameFi, nil
}
