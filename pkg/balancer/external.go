package balancer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
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

func GetStockInfo(ticker string) (stock GeographyPosition) {
	ticker = replaceAt(ticker)
	url := "https://query1.finance.yahoo.com/v10/finance/quoteSummary/" + ticker + "?modules=assetProfile"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)
	if res.StatusCode == 200 {
		stock.Country = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		if stock.Country == "United States" {
			stock.Country = "USA"
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
			stock.Country = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		}
	}
	stock.MarketType = getMarketCountry(stock.Country)
	stock.Area = GetArea(stock.Country)
	stock.Country = replaceCountry(stock.Country)
	return stock
}

// TODO: сделать нормально
func replaceCountry(name string) string {
	replaces := map[string]string{}
	replaces["USA"] = "США"
	replaces["America"] = "Америка"
	replaces["Russia"] = "Россия"
	replaces["Germany"] = "Германия"
	replaces["China"] = "Китай"
	replaces["Kazakhstan"] = "Китай"
	replaces["Europe"] = "Европа"
	replaces["Asia"] = "Азия"

	if replaces[name] != "" {
		return replaces[name]
	}
	return name
}

func replaceAt(name string) string {
	return strings.Replace(name, "@", ".", 1)
}

func getMarketCountry(name string) string {
	developedFile, err := os.Open("developed.txt")
	errorHandle(err)

	defer developedFile.Close()

	byteValue, _ := ioutil.ReadAll(developedFile)

	regexp := regexp.MustCompile(name)
	developed := regexp.FindString(string(byteValue))
	if developed != "" {
		return "Развитый"
	} else {
		return "Развивающийся"
	}
}

func GetArea(name string) string {
	jsonFile, err := os.Open("zones.json")
	errorHandle(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var areas map[string]interface{}
	err = json.Unmarshal(byteValue, &areas)
	errorHandle(err)
	for area, data := range areas {
		for _, country := range data.([]interface{}) {
			if name == country {
				return replaceCountry(area)
			}
		}
	}

	return "Неизвестно"
}
