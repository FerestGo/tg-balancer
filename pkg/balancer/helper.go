package balancer

import (
	"log"
	"sort"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
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

type kv struct {
	Key   string
	Value float64
}

type myMap map[string]float64

func mapToSortedSlice(m myMap) []kv {
	var ss []kv

	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	return ss
}
