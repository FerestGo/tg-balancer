package balancer

import t "github.com/FerestGo/investapi"

type Position struct {
	Ticker            string
	Figi              string
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

func GetPosition(receivedPosition *t.PortfolioPosition) (currentPosition Position) {
	currentPosition.Figi = receivedPosition.Figi
	currency := GetCurrency(string(receivedPosition.AveragePositionPrice.Currency))
	currentPosition.Sum = ((float64(receivedPosition.AveragePositionPrice.Units) * float64(receivedPosition.Quantity.Units)) + float64(receivedPosition.ExpectedYield.Units)) * currency

	currentPosition.GeographyPosition = GetStockInfo(receivedPosition.Figi)
	currentPosition.Currency = receivedPosition.AveragePositionPrice.Currency
	// currentPosition.Type, currentPosition.Currency = GetTypeETF(receivedPosition.Figi)

	if receivedPosition.InstrumentType == "bond" || receivedPosition.InstrumentType == "share" {
		currentPosition.Currency = receivedPosition.AveragePositionPrice.Currency
		currentPosition.Type = string(receivedPosition.InstrumentType)
	}

	if receivedPosition.InstrumentType == "Currency" {
		currentPosition.Type = "Currency"
	}

	// TODO: подставить новые figi для валют
	if receivedPosition.Figi == "USD000UTSTOM" {
		currentPosition.Currency = "USD"
		currentPosition.GeographyPosition.Country = "USA"
	}

	if receivedPosition.Figi == "EUR_RUB__TOM" {
		currentPosition.Currency = "EUR"
		currentPosition.GeographyPosition.Country = "Europe"
	}

	// TODO: FIGI вечных портфелей
	// if receivedPosition.Ticker == "TRUR" || receivedPosition.Ticker == "TEUR" || receivedPosition.Ticker == "TUSD" {
	// 	currentPosition.Type = "Balanced"
	// }

	return currentPosition
}
