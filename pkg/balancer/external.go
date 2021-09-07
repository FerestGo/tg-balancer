package balancer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type ETF struct {
	Class             string
	SubClass          string
	Currency          string
	GeographyPosition GeographyPosition `json:"Geography"`
}

func GetTypeETF(ticker string) (class string, currency string) {
	jsonFile, err := os.Open("etfs.json")
	errorHandle(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs := map[string]ETF{}
	err = json.Unmarshal(byteValue, &etfs)
	errorHandle(err)

	return etfs[ticker].Class, etfs[ticker].Currency
}

func GetGeographyETF(ticker string) GeographyPosition {
	jsonFile, err := os.Open("etfs.json")
	errorHandle(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs := map[string]ETF{}
	err = json.Unmarshal(byteValue, &etfs)
	errorHandle(err)

	return etfs[ticker].GeographyPosition
}

func GetCountryStock(ticker string) (country string) {
	ticker = replaceAt(ticker)
	url := "https://query1.finance.yahoo.com/v10/finance/quoteSummary/" + ticker + "?modules=assetProfile"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)
	if res.StatusCode == 200 {
		country = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		if country == "United States" {
			country = "USA"
		}
	} else {
		url = "https://query1.finance.yahoo.com/v10/finance/quoteSummary/" + ticker + ".ME?modules=assetProfile"

		req, _ = http.NewRequest("GET", url, nil)

		res, _ = http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ = ioutil.ReadAll(res.Body)

		var result map[string]interface{}

		json.Unmarshal([]byte(body), &result)
		if res.StatusCode == 200 {
			country = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		}
	}
	return replaceCountry(country)
}

// TODO: сделать нормально
func replaceCountry(name string) string {
	replaces := map[string]string{}
	replaces["USA"] = "США"
	replaces["Russia"] = "Россия"
	replaces["Germany"] = "Германия"
	replaces["China"] = "Китай"
	replaces["Kazakhstan"] = "Китай"

	if replaces[name] != "" {
		return replaces[name]
	}
	return name
}

func replaceAt(name string) string {
	return strings.Replace(name, "@", ".", 1)
}
