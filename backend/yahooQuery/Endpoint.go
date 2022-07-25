package yahooQuery

import (
	"backend/config"
	"backend/model"
	"backend/persistance"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func RetornaSpot() model.RetornoSpot {
	var retorno model.RetornoSpot
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
