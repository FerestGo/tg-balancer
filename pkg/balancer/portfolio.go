package balancer

import (
	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

type Portfolio struct {
	Total                 float64
	Currency              Currency
	PercentCurrency       Currency
	Geography             Geography
	PercentGeography      Geography
	StockGeography        Geography
	PercentStockGeography Geography
	Types                 Type
	PercentType           Type
	PercentTypeNoCurrency Type
}

type Type struct {
	Currency     float64
	Bonds        float64
	Stock        float64
	UndefinedEtf float64
	Goods        float64
}

type Currency struct {
	RUB float64
	USD float64
	EUR float64
}

type Geography struct {
	Area       map[string]float64
	Country    map[string]float64
	MarketType map[string]float64
}

func (portfolio *Portfolio) Analysis(receivedPortfolio sdk.Portfolio) {

	for _, position := range receivedPortfolio.Positions {
		portfolio.AddPosition(position)
	}

	return
}

func (portfolio *Portfolio) AddPosition(position sdk.PositionBalance) (receivedPortfolio *sdk.Portfolio) {
	currentPosition := GetPosition(position)

	if portfolio.StockGeography.MarketType == nil {
		portfolio.StockGeography.MarketType = map[string]float64{}
		portfolio.PercentStockGeography.MarketType = map[string]float64{}

		portfolio.StockGeography.Country = map[string]float64{}
		portfolio.PercentStockGeography.Country = map[string]float64{}

		portfolio.StockGeography.Area = map[string]float64{}
		portfolio.PercentStockGeography.Area = map[string]float64{}
	}

	SetType(currentPosition, portfolio)
	SetGeography(currentPosition, portfolio)
	AddAllWeatherEtf(currentPosition, portfolio)
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
	portfolio.PercentType.Goods = (portfolio.Types.Goods / portfolio.Total) * 100

	// Валюты всего портфеля
	portfolio.PercentCurrency.RUB = (portfolio.Currency.RUB / portfolio.Total) * 100
	portfolio.PercentCurrency.USD = (portfolio.Currency.USD / portfolio.Total) * 100
	portfolio.PercentCurrency.EUR = (portfolio.Currency.EUR / portfolio.Total) * 100

	// Страны
	for key, country := range portfolio.StockGeography.Country {
		portfolio.PercentStockGeography.Country[key] = (country / portfolio.Types.Stock) * 100
	}

	// Зоны
	for key, area := range portfolio.StockGeography.Area {
		portfolio.PercentStockGeography.Area[key] = (area / portfolio.Types.Stock) * 100
	}
	// Рынок
	for key, market := range portfolio.StockGeography.MarketType {
		portfolio.PercentStockGeography.MarketType[key] = (market / portfolio.Types.Stock) * 100
	}

	return
}

func SetGeography(currentPosition Position, portfolio *Portfolio) {
	if currentPosition.Type == "Stock" {
		if currentPosition.GeographyPosition.Country == "" {
			currentPosition.GeographyPosition = GetStockInfo(currentPosition.Ticker)
		} else {
			portfolio.StockGeography.MarketType[currentPosition.GeographyPosition.MarketType] += currentPosition.Sum
			portfolio.StockGeography.Area[currentPosition.GeographyPosition.Area] += currentPosition.Sum
			portfolio.StockGeography.Country[currentPosition.GeographyPosition.Country] += currentPosition.Sum
		}
	}
}

func SetType(currentPosition Position, portfolio *Portfolio) {

	switch currentPosition.Type {
	case "Currency":
		portfolio.Types.Currency += currentPosition.Sum
	case "Bond":
		portfolio.Types.Bonds += currentPosition.Sum
	case "Stock":
		portfolio.Types.Stock += currentPosition.Sum
	case "Balanced":
		break
	default:
		portfolio.Types.UndefinedEtf += currentPosition.Sum
	}
}

func SetCurrency(currentPosition Position, portfolio *Portfolio) {
	switch currentPosition.Currency {
	case "RUB":
		portfolio.Currency.RUB += currentPosition.Sum
	case "USD":
		portfolio.Currency.USD += currentPosition.Sum
	case "EUR":
		portfolio.Currency.EUR += currentPosition.Sum
	}
}

// исключение для тинькофф вечников
func AddAllWeatherEtf(position Position, portfolio *Portfolio) {
	if position.Type == "Balanced" {
		portfolio.Types.Bonds += position.Sum / 2
		portfolio.Types.Stock += position.Sum / 4
		portfolio.Types.Goods += position.Sum / 4
		if position.GeographyPosition.Country == "Russia" {
			portfolio.StockGeography.Area["Russia"] += position.Sum / 4
			portfolio.StockGeography.Country["Russia"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Emerging"] += position.Sum / 4
		}
		if position.GeographyPosition.Country == "USA" {
			portfolio.StockGeography.Area["America"] += position.Sum / 4
			portfolio.StockGeography.Country["USA"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Developed"] += position.Sum / 4
		}
		if position.GeographyPosition.Country == "Europe" {
			portfolio.StockGeography.Area["Europe"] += position.Sum / 4
			portfolio.StockGeography.Country["Absent"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Developed"] += position.Sum / 4
		}
	}
}
