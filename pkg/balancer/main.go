package balancer

import (
	"context"
	"regexp"
	"strings"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var client sdk.RestClient

func InitAnalysis(message string, telegramId int) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	output := initAnalysis(ctx, message)
	return output
}

func initAnalysis(ctx context.Context, message string) string {
	namePattern := regexp.MustCompile(`\s(.*)`)
	token := strings.TrimSpace(namePattern.FindString(message))

	client = *sdk.NewRestClient(token)

	accounts, err := client.Accounts(ctx)
	if err != nil {
		return "Неверный токен"
	}

	mainPortfolio, _ := client.Portfolio(ctx, accounts[0].ID)
	errorHandle(err)

	iisPortfolio, _ := client.Portfolio(ctx, accounts[1].ID)
	errorHandle(err)

	var userPortfolio Portfolio

	userPortfolio.Analysis(mainPortfolio)
	userPortfolio.Analysis(iisPortfolio)

	userPortfolio.SetPercent()

	return output(&userPortfolio)
}
