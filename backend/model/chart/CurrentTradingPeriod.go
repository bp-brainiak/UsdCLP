package chart

type CurrentTradingPeriod struct {
	Pre     CurrentTradingPeriodEntity `json:"pre"`
	Regular CurrentTradingPeriodEntity `json:"regular"`
	Post    CurrentTradingPeriodEntity `json:"post"`
}

type CurrentTrainingPeriodArray struct {
	Element []CurrentTradingPeriodEntity `json:"element"`
}
