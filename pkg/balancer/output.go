package balancer

import (
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func output(portfolio *Portfolio) string {
	printer := message.NewPrinter(language.Russian)
	output := ""
	output += printer.Sprintf("Всего средств: %.2f руб. \n\n", portfolio.Total)

	output += printer.Sprintf("Инструменты: \n")
	output += printer.Sprintf("Акции %.1f%% (%.1f р.) \n", portfolio.PercentType.Stock, portfolio.Types.Stock)
	output += printer.Sprintf("Облигации %.1f%% (%.1f р.) \n", portfolio.PercentType.Bonds, portfolio.Types.Bonds)
	output += printer.Sprintf("Валюта %.1f%% (%.1f р.) \n", portfolio.PercentType.Currency, portfolio.Types.Currency)
	output += printer.Sprintf("Биржевые товары %.1f%% (%.1f р.) \n", portfolio.PercentType.Goods, portfolio.Types.Goods)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += printer.Sprintf("Неизвестные ETF %.1f%% (%.1f р.) \n", portfolio.PercentType.UndefinedEtf, portfolio.Types.UndefinedEtf)
	}

	output += printer.Sprintf("\nВалюты: \n")
	output += printer.Sprintf("RUB %.1f%% (%.1f р.) \n", portfolio.PercentCurrency.RUB, portfolio.Currency.RUB)
	output += printer.Sprintf("USD %.1f%% (%.1f р.) \n", portfolio.PercentCurrency.USD, portfolio.Currency.USD)
	output += printer.Sprintf("EUR %.1f%% (%.1f р.) \n\n", portfolio.PercentCurrency.EUR, portfolio.Currency.EUR)

	kv := mapToSortedSlice(portfolio.StockGeography.Country)
	output += printer.Sprintf("Страны: \n")
	for i := 0; i < len(kv); i++ {
		output += printer.Sprintf("%s %.1f%% (%.1f р.) \n", kv[i].Key, portfolio.PercentStockGeography.Country[kv[i].Key], kv[i].Value)
	}

	kv = mapToSortedSlice(portfolio.StockGeography.Area)
	output += printer.Sprintf("\nЗоны: \n")
	for i := 0; i < len(kv); i++ {
		output += printer.Sprintf("%s %.1f%% (%.1f р.) \n", kv[i].Key, portfolio.PercentStockGeography.Area[kv[i].Key], kv[i].Value)
	}

	kv = mapToSortedSlice(portfolio.StockGeography.MarketType)
	output += printer.Sprintf("\nРынки: \n")
	for i := 0; i < len(kv); i++ {
		output += printer.Sprintf("%s %.1f%% (%.1f р.) \n", kv[i].Key, portfolio.PercentStockGeography.MarketType[kv[i].Key], kv[i].Value)
	}

	return output
}
