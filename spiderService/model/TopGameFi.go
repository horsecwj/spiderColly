package model

type TopCkoGameFi struct {
	ID        int    `json:"id" `
	Coin      string `json:"coin"  gorm:"primary key;"`
	Price     string `json:"price" `
	OneDay    string `json:"one_day"  gorm:"comment:'一天涨幅';"`
	OneWeek   string `json:"one_week"  gorm:"comment:'7天涨幅';"`
	DayVolume string `json:"day_volume"  gorm:"comment:'日交易';"`
	MktCap    string `json:"mkt_cap"  gorm:"comment:'市值';" `
	LastWeek  string `json:"last_week"  gorm:"comment:'7天曲线';"`
}

type TopCmkGameFi struct {
	ID        int    `json:"id `
	Coin      string `json:"coin" gorm:"primary key;"`
	Price     string `json:"price"`
	OneDay    string `json:"one_day"`
	CoinPic   string `json:"coin_pic"`
	DayVolume string `json:"day_volume"`
}

type BybitArticle struct {
	Link      string `json:"link" gorm:"unique_index;"`
	Id        uint32 `json:"id" gorm:"autoincrement;"`
	Title     string `json:"title"`
	OverView  string `json:"over_view"`
	Article   string `json:"article"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
}

type BybitNewlyArticle struct {
	Link      string `json:"link" gorm:"unique_index;"`
	Id        uint32 `json:"id" gorm:"autoincrement;"`
	Title     string `json:"title"`
	OverView  string `json:"over_view"`
	Article   string `json:"article"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
}

type SlateArticle struct {
	Link      string `json:"link" gorm:"unique_index;"`
	Id        uint32 `json:"id" gorm:"autoincrement;"`
	Title     string `json:"title"`
	OverView  string `json:"over_view"`
	Article   string `json:"article"`
	Time      string `json:"time"`
	Timestamp int64  `json:"timestamp"`
}
