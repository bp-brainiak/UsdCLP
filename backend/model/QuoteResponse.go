package model

type QuoteResponse struct {
	Result []SpotInfo `json:"result"`
	Error  string     `json:"error"`
}
