package router

import (
	"regexp"
)

type Router struct {
	routes []*Route
}

type Route struct {
	Handler   func(string, int) string
	Message   string
	IsPattern bool
}

func (r *Router) Add(message string, handler func(string, int) string, isPattern bool) {
	route := &Route{
		handler,
		message,
		isPattern,
	}
	r.routes = append(r.routes, route)
}

func (r *Router) Handle(message string, telegramId int) (response string) {
	for _, route := range r.routes {
		if route.IsPattern == true {
			if r.CheckRegexp(route.Message, message) == true {
				response = route.Handler(message, telegramId)
				return
			}
		}
		if route.Message == message {
			response = route.Handler(message, telegramId)
			return
		}
	}
	return
}

func (r *Router) Get() {
	r.Add("/start", start, false)
	r.Add("/faq", faq, false)
	// r.Add(`^t\..*\S$`, balancer.InitAnalysis, true)
}

func start(command string, telegramId int) (response string) {
	response = `
Бот предоставляет дополнительную аналитику по портфелю у брокера Тинькофф Инвестиции с анализом фондов. 

Для получения информации по всем счетам нужно отправить в диалог с ботом ваш [API Token](https://tinkoffcreditsystems.github.io/invest-openapi/auth/), счет отдельно через пробел его номер (обычно 1 - основной, 2 - ИИС)

Пример:
t.frtfnrj32sSw32s41eD23w22ed2Dxwe223DwdwwD
t.frtfnrj32sSw32s41eD23w22ed2Dxwe223DwdwwD 2
	
/faq - часто задаваемые вопросы
Баги, вопросы, предложения [сюда](https://t.me/misha_petya)`
	return response
}

func faq(command string, telegramId int) (response string) {
	response = `
*Как считает валюты?*
Акции, облигации, валюты - как они отображаются в приложении.
ETF, ПИФ - по валюте фонда т.е. для FXUS будет доллар, а не рубль
	
*Что входит в географию портфеля (страны, зоны, рынки)?*
Только акции
	
*Учитывает ли Вечный Портфель Тинькофф?*
Да. Сумма позиции / 2 к облигациям, Сумма / 4 к акциям по географии, Сумма / 4 к биржевым товарам (золото)

*Учитывает ли другие смешанные фонды?*
Сейчас нет

*Насколько это безопасно?*
Бот не хранит ваш токен, не собирает персональных данных`
	return response
}

func (r *Router) CheckRegexp(pattern string, message string) bool {
	isMatch, _ := regexp.MatchString(pattern, message)
	return isMatch

}
