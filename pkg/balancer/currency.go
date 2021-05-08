package balancer

import (
	"context"
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"time"
)

type infoCurrency struct {
	Value float64
	FIGI  string
}

var currency = map[string]infoCurrency{
	"USD": {0, "BBG0013HGFT4"},
	"EUR": {0, "BBG0013HJJ31"},
}

func GetCurrency(name string) float64 {
	if name == "RUB" {
		return 1
	}

	if currency[name].Value > 0 {
		return currency[name].Value
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	candles, err := client.Candles(ctx, time.Now().AddDate(0, 0, -1), time.Now(), sdk.CandleInterval1Day, currency[name].FIGI)
	errorHandle(err)

	if len(candles) == 0 {
		for i := 0; i > -20; i-- {
			candles, err = client.Candles(ctx, time.Now().AddDate(0, 0, i), time.Now(), sdk.CandleInterval1Day, currency[name].FIGI)
			if len(candles) > 0 {
				break
			}
		}
		errorHandle(err)
	}
	currency[name] = infoCurrency{candles[0].ClosePrice, currency[name].FIGI}

	return currency[name].Value
}
