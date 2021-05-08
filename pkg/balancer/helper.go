package balancer

import (
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
	"log"
)

func errorHandle(err error) {
	if err == nil {
		return
	}

	if tradingErr, ok := err.(sdk.TradingError); ok {
		if tradingErr.InvalidTokenSpace() {
			tradingErr.Hint = "Do you use sandbox token in production environment or vise verse?"
			log.Fatalln(tradingErr)
		}
	}

	log.Fatalln(err)
}
