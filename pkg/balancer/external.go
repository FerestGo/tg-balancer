package balancer

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

type ETF struct {
	Class    string
	SubClass string
	Currency string
	Region   string
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

func GetRegionETF(ticker string) string {
	jsonFile, err := os.Open("etfs.json")
	errorHandle(err)

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs := map[string]ETF{}
	err = json.Unmarshal(byteValue, &etfs)
	errorHandle(err)

	return etfs[ticker].Region
}

func GetRegionStock(ticker string) (region string) {
	url := "https://query1.finance.yahoo.com/v10/finance/quoteSummary/" + ticker + "?modules=assetProfile"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var result map[string]interface{}

	json.Unmarshal([]byte(body), &result)
	if res.StatusCode == 200 {
		region = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		if region == "United States" {
			region = "USA"
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
			region = result["quoteSummary"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["assetProfile"].(map[string]interface{})["country"].(string)
		}
	}
	return region
}
