package chart

import "gopkg.in/guregu/null.v3"

type Quote struct {
	High   []null.Float `json:"high"`
	Close  []null.Float `json:"close"`
	Volume []null.Float `json:"volume"`
	Open   []null.Float `json:"open"`
	Low    []null.Float `json:"low"`
}
