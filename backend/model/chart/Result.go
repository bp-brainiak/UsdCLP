package chart

type Result struct {
	Meta       Meta       `json:"meta"`
	Timestamp  []uint64   `json:"timestamp"`
	Indicators Indicators `json:"indicators"`
}
