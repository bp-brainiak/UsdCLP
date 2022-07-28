package chart

type Meta struct {
	Currency             string                         `json:"currency"`
	Symbol               string                         `json:"symbol"`
	ExhangeName          string                         `json:"exhangeName"`
	InstrumentType       string                         `json:"instrumentType"`
	firstTradeDate       int64                          `json:"firstTradeDate"`
	RegularMarketTime    int                            `json:"regularMarketTime"`
	GmtOffSet            int                            `json:"gmtoffset"`
	TimeZone             string                         `json:"timezone"`
	ExchangeTimezoneName string                         `json:"exchangeTimezoneName"`
	RegularMarketPrice   float64                        `json:"regularMarketPrice"`
	ChartPreviousClose   float64                        `json:"chartPreviousClose"`
	PreviousClose        float64                        `json:"previousClose"`
	Scale                int                            `json:"scale"`
	priceHint            int                            `json:"priceHint"`
	CurrentTradingPeriod CurrentTradingPeriod           `json:"currentTradingPeriod"`
	TradingPeriods       [][]CurrentTradingPeriodEntity `json:"tradingPeriods"`
	DataGranularity      string                         `json:"dataGranularity"`
	Range                string                         `json:"range"`
	ValidRanges          []string                       `json:"validRanges"`
}
