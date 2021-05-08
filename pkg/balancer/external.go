package balancer

import (
	"encoding/json"
	"io/ioutil"
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
