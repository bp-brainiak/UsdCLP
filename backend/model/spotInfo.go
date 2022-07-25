package model

type SpotInfo struct {
	id                                int
	Language                          string  `json:"language"`
	Region                            string  `json:"region"`
	QuoteType                         string  `json:"quoteType" gorm:"column:quoteType""`
	TypeDisp                          string  `json:"typeDisp" gorm:"column:typeDisp"`
	QuoteSourceName                   string  `json:"quoteSourceName" gorm:"column:quoteSourceName"`
	Triggerable                       bool    `json:"triggerable" gorm:"column:triggerable"`
	CustomPriceAlertConfidence        string  `json:"customPriceAlertConfidence" gorm:"column:customPriceAlertConfidence"`
	Currency                          string  `json:"currency" gorm:"column:currency"`
	ShortName                         string  `json:"shortName" gorm:"column:shortName"`
	RegularMarketChange               float64 `json:"regularMarketChange" gorm:"column:regularMarketChange"`
	RegularMarketChangePercent        float64 `json:"regularMarketChangePercent" gorm:"column:regularMarketChangePercent"`
	RegularMarketPrice                float64 `json:"regularMarketPrice" gorm:"column:regularMarketPrice"`
	RegularMarketDayHigh              float64 `json:"regularMarketDayHigh" gorm:"column:regularMarketDayHigh"`
	RegularMarketDayLow               float64 `json:"regularMarketDayLow" gorm:"column:regularMarketDayLow"`
	RegularMarketPreviousClose        float64 `json:"regularMarketPreviousClose" gorm:"column:regularMarketDayHigh"`
	Bid                               float64 `json:"bid" gorm:"column:regularMarketDayHigh"`
	Ask                               float64 `json:"ask" gorm:"column:regularMarketDayHigh"`
	RegularMarketOpen                 float64 `json:"regularMarketOpen" gorm:"column:regularMarketDayHigh"`
	FiftyTwoWeekLow                   float64 `json:"fiftyTwoWeekLow" gorm:"column:regularMarketDayHigh"`
	FiftyTwoWeekHigh                  float64 `json:"fiftyTwoWeekHigh" gorm:"column:regularMarketDayHigh"`
	FiftyDayAverage                   float64 `json:"fiftyDayAverage" gorm:"column:regularMarketDayHigh"`
	TwoHundredDayAverage              float64 `json:"twoHundredDayAverage" gorm:"column:twoHundredDayAverage"`
	MessageBoardId                    string  `json:"messageBoardId" gorm:"column:messageBoardId"`
	ExchangeTimezoneName              string  `json:"exchangeTimezoneName" gorm:"column:exchangeTimezoneName"`
	ExchangeTimezoneShortName         string  `json:"exchangeTimezoneShortName" gorm:"column:exchangeTimezoneShortName"`
	GmtOffSetMilliseconds             int64   `json:"gmtOffSetMilliseconds" gorm:"column:gmtOffSetMilliseconds"`
	Market                            string  `json:"market" gorm:"column:market"`
	EsgPopulated                      bool    `json:"esgPopulated" gorm:"column:esgPopulated"`
	Exchange                          string  `json:"exchange" gorm:"column:exchange"`
	MarketState                       string  `json:"marketState" gorm:"column:marketState"`
	Tradeable                         bool    `json:"tradeable" gorm:"column:tradeable"`
	FirstTradeDateMilliseconds        int64   `json:"firstTradeDateMilliseconds" gorm:"column:firstTradeDateMilliseconds"`
	PriceHint                         int64   `json:"priceHint" gorm:"column:priceHint"`
	RegularMarketTime                 int64   `json:"regularMarketTime" gorm:"column:regularMarketTime""`
	RegularMarketDayRange             string  `json:"regularMarketDayRange" gorm:"column:regularMarketDayRange"`
	RegularMarketVolume               float64 `json:"regularMarketVolume" gorm:"column:regularMarketVolume"`
	BidSize                           float64 `json:"bidSize" gorm:"column:bidSize"`
	AskSize                           float64 `json:"askSize" gorm:"column:askSize"`
	FullExchangeName                  string  `json:"fullExchangeName" gorm:"column:fullExchangeName"`
	AverageDailyVolume3Month          float64 `json:"averageDailyVolume3Month" gorm:"column:averageDailyVolume3Month"`
	AverageDailyVolume10Day           float64 `json:"averageDailyVolume10Day" gorm:"column:averageDailyVolume10Day"`
	FiftyTwoWeekLowChange             float64 `json:"fiftyTwoWeekLowChange" gorm:"column:fiftyTwoWeekLowChange"`
	FiftyTwoWeekLowChangePercent      float64 `json:"fiftyTwoWeekLowChangePercent" gorm:"column:fiftyTwoWeekLowChangePercent"`
	FiftyTwoWeekRange                 string  `json:"fiftyTwoWeekRange" gorm:"column:fiftyTwoWeekRange"`
	FiftyTwoWeekHighChangePercent     float64 `json:"fiftyTwoWeekHighChangePercent" gorm:"column:fiftyTwoWeekHighChangePercent"`
	FiftyDayAverageChange             float64 `json:"fiftyDayAverageChange" gorm:"column:fiftyDayAverageChange"`
	FiftyDayAverageChangePercent      float64 `json:"fiftyDayAverageChangePercent" gorm:"column:fiftyDayAverageChangePercent"`
	TwoHundredDayAverageChange        float64 `json:"twoHundredDayAverageChange" gorm:"column:twoHundredDayAverageChange"`
	TwoHundredDayAverageChangePercent float64 `json:"twoHundredDayAverageChangePercent" gorm:"column:twoHundredDayAverageChangePercent"`
	SourceInterval                    float64 `json:"sourceInterval" gorm:"column:sourceInterval"`
	ExchangeDataDelayedBy             float64 `json:"exchangeDataDelayedBy" gorm:"column:exchangeDataDelayedBy"`
	Symbol                            string  `json:"symbol" gorm:"column:symbol"`
}

func (SpotInfo) TableName() string {
	return "spot_info"
}
