package balancer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"

	t "github.com/FerestGo/investapi"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
)

type ETF struct {
	Class             string
	SubClass          string
	Currency          string
	GeographyPosition GeographyPosition `json:"Geography"`
}

var etfs = map[string]ETF{}
var stocksGeo = map[string]GeographyPosition{}
var listDeveloped []byte
var areas map[string]interface{}

func InitExternal() {
	url := "https://gist.githubusercontent.com/FerestGo/67e28d85ad70ca0f531f02b64c438c80/raw/7618acd1a307c1abdd04f82c0e14cee895a3c157/etfs.json"

	req, _ := http.NewRequest("GET", url, nil)

	res, _ := http.DefaultClient.Do(req)
	jsonFile := res.Body

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	etfs = map[string]ETF{}
	err := json.Unmarshal(byteValue, &etfs)
	errorHandle(err)
	url = "https://gist.githubusercontent.com/FerestGo/97993663b544397735458d85b14049a3/raw/37ac956bd14c146395864cc3b63542ba9cdd0c86/developed.txt"

	req, _ = http.NewRequest("GET", url, nil)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Cant get market country %s:", err.Error())
	}
	jsonFile = res.Body

	defer jsonFile.Close()

	listDeveloped, _ = ioutil.ReadAll(jsonFile)

	url = "https://gist.githubusercontent.com/FerestGo/c90b19f17bd6f41ce92a4b9c16b14f21/raw/1926a54f24f406f2441e99eb6fc7055a2e936650/zones.json"

	req, _ = http.NewRequest("GET", url, nil)

	res, err = http.DefaultClient.Do(req)
	if err != nil {
		fmt.Errorf("Cant get area country %s:", err.Error())
	}
	jsonFile = res.Body
	byteValue, _ = ioutil.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &areas)
	errorHandle(err)
}

func GetTypeETF(ticker string) (class string, currency string) {
	return etfs[ticker].Class, etfs[ticker].Currency
}

func GetGeographyETF(ticker string) GeographyPosition {
	return etfs[ticker].GeographyPosition
}

func GetStockInfo(figi string) (stock GeographyPosition) {
	conn, err := grpc.Dial("invest-public-api.tinkoff.ru:443",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: token,
		})))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	insturement, _ := t.NewInstrumentsServiceClient(conn).GetInstrumentBy(context.Background(), &t.InstrumentRequest{
		IdType:    t.InstrumentIdType_INSTRUMENT_ID_TYPE_FIGI,
		ClassCode: "",
		Id:        figi,
	})
	// TODO обработка ошибки
	stock.Country = insturement.Instrument.CountryOfRiskName
	stock.MarketType = getMarketCountry(stock.Country)
	stock.Area = GetArea(stock.Country)
	stock.Country = replaceCountry(stock.Country)
	stocksGeo[insturement.Instrument.Ticker] = stock
	return stock
}

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
	regexp := regexp.MustCompile(name)
	developed := regexp.FindString(string(listDeveloped))
	if developed != "" {
		return "Развитый"
	} else {
		return "Развивающийся"
	}
}

func GetArea(name string) string {
	for area, data := range areas {
		for _, country := range data.([]interface{}) {
			if name == country {
				return replaceCountry(area)
			}
		}
	}

	return "Неизвестно"
}
