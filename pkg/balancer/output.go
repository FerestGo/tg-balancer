package balancer

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func output(portfolio *Portfolio) string {
	printer := message.NewPrinter(language.Russian)
	output := ""
	output += printer.Sprintf("Всего средств: %.2f руб. \n\n", portfolio.Total)
	output += printer.Sprintf("Валюты: \n")
	output += printer.Sprintf("RUB %.0f%% (%.0f р.) \n", portfolio.PercentCurrency.RUB, portfolio.Currency.RUB)
	output += printer.Sprintf("USD %.0f%% (%.0f р.) \n", portfolio.PercentCurrency.USD, portfolio.Currency.USD)
	output += printer.Sprintf("EUR %.0f%% (%.0f р.) \n\n", portfolio.PercentCurrency.EUR, portfolio.Currency.EUR)

	output += printer.Sprintf("Инструменты: \n")
	output += printer.Sprintf("Акции %.0f%% (%.0f р.) \n", portfolio.PercentType.Stock, portfolio.Types.Stock)
	output += printer.Sprintf("Облигации %.0f%% (%.0f р.) \n", portfolio.PercentType.Bonds, portfolio.Types.Bonds)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += printer.Sprintf("Неизвестные ETF %.0f%% (%.0f р.) \n", portfolio.PercentType.UndefinedEtf, portfolio.Types.UndefinedEtf)

	}
	output += printer.Sprintf("Валюта %.0f%% (%.0f р.) \n", portfolio.PercentType.Currency, portfolio.Types.Currency)
	output += printer.Sprintf("Биржевые товары %.0f%% (%.0f р.) \n\n", portfolio.PercentType.Goods, portfolio.Types.Goods)

	output += printer.Sprintf("Инструменты без валюты: \n")
	output += printer.Sprintf("Акции %.0f%% (%.0f р.) \n", portfolio.PercentTypeNoCurrency.Stock, portfolio.Types.Stock)
	output += printer.Sprintf("Облигации %.0f%% (%.0f р.) \n", portfolio.PercentTypeNoCurrency.Bonds, portfolio.Types.Bonds)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += printer.Sprintf("Неизвестные ETF %.0f%% (%.0f р.) \n\n", portfolio.PercentTypeNoCurrency.UndefinedEtf, portfolio.Types.UndefinedEtf)

	}
	output += printer.Sprintf("Биржевые товары %.0f%% (%.0f р.) \n", portfolio.PercentTypeNoCurrency.Goods, portfolio.Types.Goods)

	output += printer.Sprintf("\nГеография акций (experimentally): \n")
	output += printer.Sprintf("\nСтраны: \n")
	for key, _ := range portfolio.StockGeography.Country {
		output += printer.Sprintf("%s %.0f%% (%.0f р.) \n", key, portfolio.PercentStockGeography.Country[key], portfolio.StockGeography.Country[key])
	}

	output += printer.Sprintf("\nЗоны: \n")
	for key, _ := range portfolio.StockGeography.Area {
		output += printer.Sprintf("%s %.0f%% (%.0f р.) \n", key, portfolio.PercentStockGeography.Area[key], portfolio.StockGeography.Area[key])
	}

	output += printer.Sprintf("\nРынки: \n")
	for key, _ := range portfolio.StockGeography.MarketType {
		output += printer.Sprintf("%s %.0f%% (%.0f р.) \n", key, portfolio.PercentStockGeography.MarketType[key], portfolio.StockGeography.MarketType[key])
	}

	return output
}
