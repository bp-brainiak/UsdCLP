package persistance

import (
	"backend/model/ChartPersistance"
	"backend/model/chart"
	"encoding/json"
	"fmt"
	"time"
)

func Create(chartData chart.ChartModel, requestTime time.Time, _range string, _region string, _interval string, _lang string) {

	data, _ := json.Marshal(chartData)
	chartToPersist := &ChartPersistance.Chart{
		Range:       _range,
		RequestTime: requestTime,
		Region:      _region,
		Interval:    _interval,
		Lang:        _lang,
		ResultData:  fmt.Sprintf("%s", data),
	}
	db := GetDatabase()
	db.Create(&chartToPersist)

}
