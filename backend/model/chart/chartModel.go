package chart

import "backend/model"

type ChartModel struct {
	Chart Chart `json:"chart"`
}

type Chart struct {
	Result []Result    `json:"result"`
	Error  model.Error `json:"error"`
}
