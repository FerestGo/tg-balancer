package balancer

import (
	"fmt"
)

func output(portfolio *Portfolio) string {
	output := ""
	output += fmt.Sprintf("Всего средств: %.2f руб. \n\n", portfolio.Total)
	output += fmt.Sprintf("Валюты: \n")
	output += fmt.Sprintf("RUB %.2f%% \n", portfolio.PercentCurrency.RUB)
	output += fmt.Sprintf("USD %.2f%% \n", portfolio.PercentCurrency.USD)
	output += fmt.Sprintf("EUR %.2f%% \n\n", portfolio.PercentCurrency.EUR)

	output += fmt.Sprintf("Инструменты: \n")
	output += fmt.Sprintf("Акции %.2f%% \n", portfolio.PercentType.Stock)
	output += fmt.Sprintf("Облигации %.2f%% \n", portfolio.PercentType.Bonds)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += fmt.Sprintf("Неизвестные ETF %.2f%% \n", portfolio.PercentType.UndefinedEtf)

	}
	output += fmt.Sprintf("Валюта %.2f%% \n", portfolio.PercentType.Currency)
	output += fmt.Sprintf("Золото %.2f%% \n\n", portfolio.PercentType.Gold)

	output += fmt.Sprintf("Инструменты без валют: \n")
	output += fmt.Sprintf("Акции %.2f%% \n", portfolio.PercentTypeNoCurrency.Stock)
	output += fmt.Sprintf("Облигации %.2f%% \n", portfolio.PercentTypeNoCurrency.Bonds)
	if portfolio.PercentType.UndefinedEtf > 0 {
		output += fmt.Sprintf("Неизвестные ETF %.2f%% \n\n", portfolio.PercentTypeNoCurrency.UndefinedEtf)

	}
	output += fmt.Sprintf("Золото %.2f%% \n\n", portfolio.PercentTypeNoCurrency.Gold)

	output += fmt.Sprintf("Акции по регионам: \n")
	output += fmt.Sprintf("США %.2f%% \n", portfolio.StockPercentRegion.USA)
	output += fmt.Sprintf("Россия %.2f%% \n", portfolio.StockPercentRegion.RU)
	output += fmt.Sprintf("Европа %.2f%% \n", portfolio.StockPercentRegion.Europe)
	output += fmt.Sprintf("Китай %.2f%% \n", portfolio.StockPercentRegion.China)
	output += fmt.Sprintf("Мир %.2f%% \n", portfolio.StockPercentRegion.World)
	//output += fmt.Sprintf("Развивающиеся %.2f%% \n", portfolio.StockPercentRegion.Developing)
	if portfolio.StockPercentRegion.Undefined > 0 {
		output += fmt.Sprintf("Неопределенные %.2f%% \n\n", portfolio.StockPercentRegion.Undefined)
	}

	output += fmt.Sprintf("Весь портфель по регионам: \n")
	output += fmt.Sprintf("США %.2f%% \n", portfolio.PercentRegion.USA)
	output += fmt.Sprintf("Россия %.2f%% \n", portfolio.PercentRegion.RU)
	output += fmt.Sprintf("Европа %.2f%% \n", portfolio.PercentRegion.Europe)
	output += fmt.Sprintf("Китай %.2f%% \n", portfolio.PercentRegion.China)
	output += fmt.Sprintf("Мир %.2f%% \n", portfolio.PercentRegion.World)
	//output += fmt.Sprintf("Развивающиеся %.2f%% \n", portfolio.PercentRegion.Developing)
	if portfolio.PercentRegion.Undefined > 0 {
		output += fmt.Sprintf("Неопределенные %.2f%%", portfolio.PercentRegion.Undefined)
	}
	return output
}
