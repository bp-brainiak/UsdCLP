package spot

import "backend/model"

type QuoteResponse struct {
	Result []SpotInfo  `json:"result"`
	Error  model.Error `json:"error"`
}
