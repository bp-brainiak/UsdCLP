package chart

type CurrentTradingPeriodEntity struct {
	TimeZone  string `json:"timezone"`
	Start     uint64 `json:"start"`
	End       uint64 `json:"end"`
	GmtOffset int    `json:"gmtoffset"`
}
