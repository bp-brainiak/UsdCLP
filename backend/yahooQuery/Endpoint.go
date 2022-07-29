package yahooQuery

import (
	"backend/config"
	"backend/model/chart"
	"backend/model/spot"
	"backend/persistance"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func RetornaSpot() spot.RetornoSpot {
	var retorno spot.RetornoSpot
	var endpointConfig = config.GetValue("dolarSpot")
	var keyRest = config.GetValue("yhKey")
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointConfig, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-API-KEY", keyRest)
	req.Header.Add("Content-Type", "application/json")
	resp, err2 := client.Do(req)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(body, &retorno)
	if err != nil {
		log.Fatal(err)
	}
	persistance.PersistDolarInfo(retorno.QuoteResponse.Result[0])
	return retorno
}

func GetChart(_range string, _region string, _interval string, _lang string) chart.ChartModel {
	var retorno chart.ChartModel
	var retornoP chart.ChartModel
	var endpointUrl = fmt.Sprintf("%srange=%s&region=%s&interval=%s&lang=%s",
		config.GetValue("dolarChart"),
		_range,
		_region,
		_interval,
		_lang,
	)

	fmt.Println(endpointUrl)
	var keyRest = config.GetValue("yhKey")
	client := &http.Client{}
	req, err := http.NewRequest("GET", endpointUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Add("X-API-KEY", keyRest)
	req.Header.Add("Content-Type", "application/json")
	t := time.Now()
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(body, &retorno)
	if err != nil {
		log.Fatal(err)
	}
	retornoP = retorno
	persistance.Create(retornoP, t, _range, _region, _interval, _lang)
	return retorno

}
