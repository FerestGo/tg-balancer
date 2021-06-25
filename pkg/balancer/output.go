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
	output += printer.Sprintf("Золото %.0f%% (%.0f р.) \n\n", portfolio.PercentType.Gold, portfolio.Types.Gold)

	output += printer.Sprintf("Инструменты без валюты: \n")
	output += printer.Sprintf("Акции %.0f%% (%.0f р.) \n", portfolio.PercentTypeNoCurrency.Stock, portfolio.Types.Stock)
	output += printer.Sprintf("Облигации %.0f%% (%.0f р.) \n", portfolio.PercentTypeNoCurrency.Bonds, portfolio.Types.Bonds)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += printer.Sprintf("Неизвестные ETF %.0f%% (%.0f р.) \n\n", portfolio.PercentTypeNoCurrency.UndefinedEtf, portfolio.Types.UndefinedEtf)

	}
	output += printer.Sprintf("Золото %.0f%% (%.0f р.) \n\n", portfolio.PercentTypeNoCurrency.Gold, portfolio.Types.Gold)

	output += printer.Sprintf("Акции по регионам: \n")
	output += printer.Sprintf("США %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.USA, portfolio.StockRegion.USA)
	output += printer.Sprintf("Россия %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.RU, portfolio.StockRegion.RU)
	output += printer.Sprintf("Европа %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.Europe, portfolio.StockRegion.Europe)
	output += printer.Sprintf("Китай %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.China, portfolio.StockRegion.China)
	output += printer.Sprintf("Мир %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.World, portfolio.StockRegion.World)
	//output += printer.Sprintf("Развивающиеся %.0f%% (%.0f р.) \n", portfolio.StockPercentRegion.Developing)
	if portfolio.StockPercentRegion.Undefined > 0 {
		output += printer.Sprintf("Неопределенные %.0f%% (%.0f р.) \n\n", portfolio.StockPercentRegion.Undefined, portfolio.StockRegion.Undefined)
	}

	output += printer.Sprintf("Весь портфель по регионам: \n")
	output += printer.Sprintf("США %.0f%% (%.0f р.) \n", portfolio.PercentRegion.USA, portfolio.Region.USA)
	output += printer.Sprintf("Россия %.0f%% (%.0f р.) \n", portfolio.PercentRegion.RU, portfolio.Region.RU)
	output += printer.Sprintf("Европа %.0f%% (%.0f р.) \n", portfolio.PercentRegion.Europe, portfolio.Region.Europe)
	output += printer.Sprintf("Китай %.0f%% (%.0f р.) \n", portfolio.PercentRegion.China, portfolio.Region.China)
	output += printer.Sprintf("Мир %.0f%% (%.0f р.) \n", portfolio.PercentRegion.World, portfolio.Region.World)
	//output += printer.Sprintf("Развивающиеся %.0f%% (%.0f р.) \n", portfolio.PercentRegion.Developing)
	if portfolio.PercentRegion.Undefined > 0 {
		output += printer.Sprintf("Неопределенные %.0f%%", portfolio.PercentRegion.Undefined)
	}
	return output
}
