package balancer

import (
	t "github.com/FerestGo/investapi"
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

func (portfolio *Portfolio) Analysis(receivedPortfolio *t.PortfolioResponse) {

	for _, position := range receivedPortfolio.Positions {
		portfolio.AddPosition(position)
	}

	// TODO: обработка валют если нужна
	// for _, currency := range receivedPortfolio.Currencies {
	// 	portfolio.AddCurrency(currency)
	// }

	// if receivedPortfolio.Positions[0].Ticker == "AAPL" {
	//  test := receivedPortfolio.Positions[0]
	//  test.Balance = 50
	//  portfolio.AddPosition(test)
	// }
	return
}

func (portfolio *Portfolio) AddPosition(position *t.PortfolioPosition) {
	currentPosition := GetPosition(position)

	// TODO: вспомнить зачем это
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

func (portfolio *Portfolio) AddCurrency(currency sdk.CurrencyBalance) (receivedPortfolio *sdk.Portfolio) {

	if currency.Currency == "rub" {
		portfolio.Currency.RUB += currency.Balance
		portfolio.Types.Currency += currency.Balance
		portfolio.Total += currency.Balance
	}

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
	if currentPosition.Type == "share" {
		if currentPosition.GeographyPosition.Country == "" {
			currentPosition.GeographyPosition = GetStockInfo(currentPosition.Ticker)
		}
		portfolio.StockGeography.MarketType[currentPosition.GeographyPosition.MarketType] += currentPosition.Sum
		portfolio.StockGeography.Area[currentPosition.GeographyPosition.Area] += currentPosition.Sum
		portfolio.StockGeography.Country[currentPosition.GeographyPosition.Country] += currentPosition.Sum
	}
}

func SetType(currentPosition Position, portfolio *Portfolio) {

	switch currentPosition.Type {
	case "currency":
		portfolio.Types.Currency += currentPosition.Sum
	case "bond":
		portfolio.Types.Bonds += currentPosition.Sum
	case "share":
		portfolio.Types.Stock += currentPosition.Sum
	case "balanced":
		break
	default:
		portfolio.Types.UndefinedEtf += currentPosition.Sum
	}
}

func SetCurrency(currentPosition Position, portfolio *Portfolio) {
	switch currentPosition.Currency {
	case "rub":
		portfolio.Currency.RUB += currentPosition.Sum
	case "usd":
		portfolio.Currency.USD += currentPosition.Sum
	case "eur":
		portfolio.Currency.EUR += currentPosition.Sum
	}
}

// исключение для тинькофф вечников
func AddAllWeatherEtf(position Position, portfolio *Portfolio) {
	if position.Type == "Balanced" {
		portfolio.Types.Bonds += position.Sum / 2
		portfolio.Types.Stock += position.Sum / 4
		portfolio.Types.Goods += position.Sum / 4
		if position.GeographyPosition.Country == "Россия" {
			portfolio.StockGeography.Area["Россия"] += position.Sum / 4
			portfolio.StockGeography.Country["Россия"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Развивающийся"] += position.Sum / 4
		}
		if position.GeographyPosition.Country == "США" {
			portfolio.StockGeography.Area["Америка"] += position.Sum / 4
			portfolio.StockGeography.Country["США"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Развитый"] += position.Sum / 4
		}
		if position.GeographyPosition.Country == "Европа" {
			portfolio.StockGeography.Area["Европа"] += position.Sum / 4
			portfolio.StockGeography.Country["Мир"] += position.Sum / 4
			portfolio.StockGeography.MarketType["Развитый"] += position.Sum / 4
		}
	}
}
