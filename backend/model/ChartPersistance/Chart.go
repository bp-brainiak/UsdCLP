package ChartPersistance

import "time"

type Chart struct {
	Id          int       `gorm:"primaryKey" gorm:"<-:false"`
	Range       string    `gorm:"column:range"`
	Region      string    `gorm:"column:region"`
	Interval    string    `gorm:"column:interval"`
	Lang        string    `gorm:"column:lang"`
	RequestTime time.Time `gorm:"column:RequestTime"`
	ResultData  string    `gorm:"column:resultData"`
}

func (Chart) TableName() string {
	return "Chart"
}
