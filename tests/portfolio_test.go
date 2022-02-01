package tests

import (
	"fmt"
	"testing"

	"github.com/FerestGo/tg-balancer/pkg/balancer"
	"github.com/FerestGo/tg-balancer/pkg/config"
)

// var jsonPortfolio = []byte(`{"Positions":[{"figi":"BBG000B9XRY4","ticker":"AAPL","isin":"US0378331005","instrumentType":"Stock","balance":1,"blocked":0,"lots":1,"expectedYield":{"currency":"USD","value":38.32},"averagePositionPrice":{"currency":"USD","value":111.51},"averagePositionPriceNoNkd":{"currency":"","value":0},"name":"Apple"},{"figi":"BBG333333333","ticker":"TMOS","isin":"RU000A101X76","instrumentType":"Etf","balance":11018,"blocked":0,"lots":11018,"expectedYield":{"currency":"RUB","value":9178.98},"averagePositionPrice":{"currency":"RUB","value":6.55},"averagePositionPriceNoNkd":{"currency":"","value":0},"name":"Тинькофф iMOEX"},{"figi":"TCS00A102EQ8","ticker":"TSPX","isin":"RU000A102EQ8","instrumentType":"Etf","balance":14600,"blocked":0,"lots":146,"expectedYield":{"currency":"USD","value":156.22},"averagePositionPrice":{"currency":"USD","value":0.1075},"averagePositionPriceNoNkd":{"currency":"","value":0},"name":"Тинькофф S&P 500"},{"figi":"BBG0013HGFT4","ticker":"USD000UTSTOM","isin":"","instrumentType":"Currency","balance":0.67,"blocked":0,"lots":0,"expectedYield":{"currency":"RUB","value":-0.06},"averagePositionPrice":{"currency":"RUB","value":71.595},"averagePositionPriceNoNkd":{"currency":"","value":0},"name":"Доллар США"}],"Currencies":[{"currency":"RUB","balance":5.32,"blocked":0},{"currency":"USD","balance":0.67,"blocked":0}]}`)
// var sdkPortfolio tinksdk.Portfolio

func TestBalancer(t *testing.T) {
	cfg, err := config.Init()
	if err != nil {
		fmt.Printf("Config init error: %s", err.Error())
	}
	fmt.Println(balancer.InitAnalysis(cfg.T_Token, 71783442))
}
