package util

import (
	"Spider/spiderService/model"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetTopGameFiCoinCko() ([]*model.TopCkoGameFi, error) {
	var ArrTopGameFi = make([]*model.TopCkoGameFi, 0, 30)
	var err error
	c := colly.NewCollector(
		colly.MaxDepth(7),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:96.0) Gecko/20100101 Firefox/96.0"),
	)
	// 设置抓取频率限制
	_ = c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		RandomDelay: 3 * time.Second, // 随机延迟
	})
	c.OnRequest(func(req *colly.Request) {
		log.Println("Visiting", req.URL)
	})
	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong:", err)
	})

	c.OnHTML("table tbody tr", func(elem *colly.HTMLElement) {
		elem.DOM.Each(func(_ int, s *goquery.Selection) {

			str := s.Find("td")
			link, _ := str.Eq(10).Find("img").Attr("data-src")
			res, _ := http.Get(link)
			data, _ := ioutil.ReadAll(res.Body)

			//base64Data := "data:image/svg+xml;base64," + base64.StdEncoding.EncodeToString(data)
			//log.Println(base64Data)

			tplData := model.TopCkoGameFi{
				ID:        len(ArrTopGameFi),
				Coin:      strings.ReplaceAll(str.Eq(2).Find("a").Eq(1).Text(), "\n", ""),
				GameFi:    strings.ReplaceAll(str.Eq(2).Find("a").Eq(0).Text(), "\n", ""),
				Price:     str.Eq(4).Find("span").Text(),
				OneDay:    str.Eq(6).Find("span").Text(),
				OneWeek:   str.Eq(7).Find("span").Text(),
				DayVolume: str.Eq(8).Find("span").Text(),
				MktCap:    str.Eq(9).Find("span").Text(),
				LastWeek:  string(data), //svg+xml文件
			}
			ArrTopGameFi = append(ArrTopGameFi, &tplData)

		})
	})
	// 查找下一页
	c.OnHTML("li[class='page-item next'] a", func(element *colly.HTMLElement) {
		href, found := element.DOM.Attr("href")
		// 如果有下一页，则继续访问
		if found && href != "#" {
			_ = element.Request.Visit(element.Request.AbsoluteURL("https://www.coingecko.com" + href))
		}
	})

	err = c.Visit("https://www.coingecko.com/en/categories/gaming?page=1")
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	c.Wait()
	return ArrTopGameFi, nil
}
