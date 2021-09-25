package balancer

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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
	url := "https://gist.githubusercontent.com/FerestGo/67e28d85ad70ca0f531f02b64c438c80/raw/7618acd1a307c1abdd04f82c0e14cee895a3c157/etfs.json"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	jsonFile := res.Body

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs := map[string]ETF{}
	err := json.Unmarshal(byteValue, &etfs)
	errorHandle(err)

	return etfs[ticker].Class, etfs[ticker].Currency
}

func GetGeographyETF(ticker string) GeographyPosition {
	url := "https://gist.githubusercontent.com/FerestGo/67e28d85ad70ca0f531f02b64c438c80/raw/7618acd1a307c1abdd04f82c0e14cee895a3c157/etfs.json"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	jsonFile := res.Body

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs := map[string]ETF{}
	err := json.Unmarshal(byteValue, &etfs)
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
	url := "https://gist.githubusercontent.com/FerestGo/97993663b544397735458d85b14049a3/raw/37ac956bd14c146395864cc3b63542ba9cdd0c86/developed.txt"

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Cant get market country %s:", err.Error())
	}
	jsonFile := res.Body

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	regexp := regexp.MustCompile(name)
	developed := regexp.FindString(string(byteValue))
	if developed != "" {
		return "Развитый"
	} else {
		return "Развивающийся"
	}
}

func GetArea(name string) string {
	url := "https://gist.githubusercontent.com/FerestGo/c90b19f17bd6f41ce92a4b9c16b14f21/raw/1926a54f24f406f2441e99eb6fc7055a2e936650/zones.json"

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Cant get area country %s:", err.Error())
	}
	jsonFile := res.Body
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
