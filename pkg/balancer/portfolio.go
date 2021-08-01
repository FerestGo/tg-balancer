package balancer

import (
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

type Portfolio struct {
	Total                 float64
	Currency              Currency
	PercentCurrency       Currency
	Region                Region
	PercentRegion         Region
	StockPercentRegion    Region
	StockRegion           Region
	Types                 Type
	PercentType           Type
	PercentTypeNoCurrency Type
}

type Type struct {
	Currency     float64
	Bonds        float64
	Stock        float64
	UndefinedEtf float64
	Gold         float64
}

type Currency struct {
	RUB float64
	USD float64
	EUR float64
}

type Region struct {
	World      float64
	USA        float64
	RU         float64
	Europe     float64
	Developing float64
	China      float64
	Undefined  float64
}

func (portfolio *Portfolio) Analysis(receivedPortfolio sdk.Portfolio) {
	for _, position := range receivedPortfolio.Positions {
		portfolio.SetParamsByName(position)
	}

	return
}

func (portfolio *Portfolio) SetParamsByName(position sdk.PositionBalance) (receivedPortfolio *sdk.Portfolio) {
	currentPosition := GetPosition(position)

	SetRegion(currentPosition, portfolio)
	AddAllWeatherEtf(currentPosition, portfolio)
	SetType(currentPosition, portfolio)
	SetCurrency(currentPosition, portfolio)

	portfolio.Total += currentPosition.Sum

	return
}

func (portfolio *Portfolio) SetPercent() {
	// Типы
	portfolio.PercentType.Currency = (portfolio.Types.Currency / portfolio.Total) * 100
	portfolio.PercentType.Stock = (portfolio.Types.Stock / portfolio.Total) * 100
	portfolio.PercentType.Bonds = (portfolio.Types.Bonds / portfolio.Total) * 100
	portfolio.PercentType.UndefinedEtf = (portfolio.Types.UndefinedEtf / portfolio.Total) * 100
	portfolio.PercentType.Gold = (portfolio.Types.Gold / portfolio.Total) * 100

	// Типы без валют
	portfolio.PercentTypeNoCurrency.Stock = (portfolio.Types.Stock / (portfolio.Total - portfolio.Types.Currency)) * 100
	portfolio.PercentTypeNoCurrency.Bonds = (portfolio.Types.Bonds / (portfolio.Total - portfolio.Types.Currency)) * 100
	portfolio.PercentTypeNoCurrency.UndefinedEtf = (portfolio.Types.UndefinedEtf / (portfolio.Total - portfolio.Types.Currency)) * 100
	portfolio.PercentTypeNoCurrency.Gold = (portfolio.Types.Gold / (portfolio.Total - portfolio.Types.Currency)) * 100

	// Валюты всего портфеля
	portfolio.PercentCurrency.RUB = (portfolio.Currency.RUB / portfolio.Total) * 100
	portfolio.PercentCurrency.USD = (portfolio.Currency.USD / portfolio.Total) * 100
	portfolio.PercentCurrency.EUR = (portfolio.Currency.EUR / portfolio.Total) * 100

	// Регионы акций
	portfolio.StockPercentRegion.RU = (portfolio.StockRegion.RU / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.Europe = (portfolio.StockRegion.Europe / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.USA = (portfolio.StockRegion.USA / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.China = (portfolio.StockRegion.China / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.Developing = (portfolio.StockRegion.Developing / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.World = (portfolio.StockRegion.World / portfolio.Types.Stock) * 100
	portfolio.StockPercentRegion.Undefined = (portfolio.StockRegion.Undefined / portfolio.Types.Stock) * 100

	// Регионы
	portfolio.PercentRegion.RU = (portfolio.Region.RU / portfolio.Total) * 100
	portfolio.PercentRegion.Europe = (portfolio.Region.Europe / portfolio.Total) * 100
	portfolio.PercentRegion.USA = (portfolio.Region.USA / portfolio.Total) * 100
	portfolio.PercentRegion.China = (portfolio.Region.China / portfolio.Total) * 100
	portfolio.PercentRegion.Developing = (portfolio.Region.Developing / portfolio.Total) * 100
	portfolio.PercentRegion.World = (portfolio.Region.World / portfolio.Total) * 100
	portfolio.PercentRegion.Undefined = (portfolio.Region.Undefined / portfolio.Total) * 100

	return
}

func AddAllWeatherEtf(position Position, portfolio *Portfolio) {
	if position.Type == "Balanced" {
		// исключение для тинькофф вечников, стоит убрать хардкод
		portfolio.Types.Bonds += position.Sum / 2
		portfolio.Types.Stock += position.Sum / 4
		portfolio.Types.Gold += position.Sum / 4
		if position.Region == "Russia" {
			portfolio.StockRegion.RU += position.Sum / 4
		}
		if position.Region == "USA" {
			portfolio.StockRegion.USA += position.Sum / 4
		}
		if position.Region == "Europe" {
			portfolio.StockRegion.Europe += position.Sum / 4
		}
	}
}

func SetRegion(currentPosition Position, portfolio *Portfolio) {
	if currentPosition.Region == "" {
		currentPosition.Region = GetRegionStock(currentPosition.Ticker)
	}
	switch currentPosition.Region {
	case "Russia":
		portfolio.Region.RU += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.RU += currentPosition.Sum
		}
	case "Europe":
		portfolio.Region.Europe += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.Europe += currentPosition.Sum
		}
	case "USA":
		portfolio.Region.USA += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.USA += currentPosition.Sum
		}
	case "World":
		portfolio.Region.World += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.World += currentPosition.Sum
		}
	case "China":
		portfolio.Region.China += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.China += currentPosition.Sum
		}
	case "Kazakhstan":
		portfolio.Region.World += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.World += currentPosition.Sum
		}
	case "Developing":
		portfolio.Region.Developing += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.Developing += currentPosition.Sum
		}
	default:
		portfolio.Region.Undefined += currentPosition.Sum
		if currentPosition.Type == "Stock" {
			portfolio.StockRegion.Undefined += currentPosition.Sum
		}
	}

}

func SetType(currentPosition Position, portfolio *Portfolio) {
	if currentPosition.Type == "Currency" {
		portfolio.Types.Currency += currentPosition.Sum
	}
	if currentPosition.Type == "Bond" {
		portfolio.Types.Bonds += currentPosition.Sum
	}
	if currentPosition.Type == "Stock" {
		portfolio.Types.Stock += currentPosition.Sum
	}
	if currentPosition.Type == "" {
		portfolio.Types.UndefinedEtf += currentPosition.Sum
	}
}

func SetCurrency(currentPosition Position, portfolio *Portfolio) {
	if currentPosition.Currency == "RUB" {
		portfolio.Currency.RUB += currentPosition.Sum
	}
	if currentPosition.Currency == "USD" {
		portfolio.Currency.USD += currentPosition.Sum
	}
	if currentPosition.Currency == "EUR" {
		portfolio.Currency.EUR += currentPosition.Sum
	}
}
