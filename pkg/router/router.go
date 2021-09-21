package router

import (
	"regexp"

	"github.com/FerestGo/tg-balancer/pkg/balancer"
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
	r.Add(`^t\.`, balancer.InitAnalysis, true)
}

func start(command string, telegramId int) (response string) {
	response = "Это бот"
	return response
}

func (r *Router) CheckRegexp(pattern string, message string) bool {
	isMatch, _ := regexp.MatchString(pattern, message)
	return isMatch

}
