package balancer

import sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"

type Position struct {
	Ticker            string
	Currency          string
	Type              string
	GeographyPosition GeographyPosition `json:"Geography"`
	Sum               float64
}

type GeographyPosition struct {
	Area       string `json:"area"`
	Country    string `json:"country"`
	MarketType string `json:"marketType"`
}

func GetPosition(receivedPosition sdk.PositionBalance) (currentPosition Position) {
	currentPosition.Ticker = receivedPosition.Ticker
	currency := GetCurrency(string(receivedPosition.AveragePositionPrice.Currency))
	currentPosition.Sum = ((receivedPosition.AveragePositionPrice.Value * receivedPosition.Balance) + receivedPosition.ExpectedYield.Value) * currency

	currentPosition.GeographyPosition = GetGeographyETF(receivedPosition.Ticker)
	currentPosition.Type, currentPosition.Currency = GetTypeETF(receivedPosition.Ticker)

	if receivedPosition.InstrumentType == "Bond" || receivedPosition.InstrumentType == "Stock" {
		currentPosition.Currency = string(receivedPosition.ExpectedYield.Currency)
		currentPosition.Type = string(receivedPosition.InstrumentType)
	}

	if receivedPosition.InstrumentType == "Currency" {
		currentPosition.Type = "Currency"
	}

	if receivedPosition.Ticker == "USD000UTSTOM" {
		currentPosition.Currency = "USD"
		currentPosition.GeographyPosition.Country = "USA"
	}

	if receivedPosition.Ticker == "EUR_RUB__TOM" {
		currentPosition.Currency = "EUR"
		currentPosition.GeographyPosition.Country = "Europe"
	}

	if receivedPosition.Ticker == "TRUR" || receivedPosition.Ticker == "TEUR" || receivedPosition.Ticker == "TUSD" {
		currentPosition.Type = "Balanced"
	}

	return currentPosition
}
