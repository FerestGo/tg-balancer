package balancer

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	sdk "github.com/TinkoffCreditSystems/invest-openapi-go-sdk"
)

var client sdk.RestClient

func InitAnalysis(message string, telegramId int) string {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	token, accountId := getAccount(message)
	output := initAnalysis(ctx, token, accountId)
	return output
}

func getAccount(message string) (token string, account int) {
	accountPattern := regexp.MustCompile(`\s\d$`)
	account, _ = strconv.Atoi(strings.TrimSpace(accountPattern.FindString(message)))
	if account != 0 {
		token = message[0 : len(message)-2]
	} else {
		token = message
	}
	return
}

func initAnalysis(ctx context.Context, token string, accountId int) string {
	mg := ""
	client = *sdk.NewRestClient(token)

	accounts, err := client.Accounts(ctx)
	if err != nil {
		return "Неверный токен"
	}
	var userPortfolio Portfolio
	var portfolio sdk.Portfolio

	if accountId != 0 && accounts[accountId-1].ID != "" {
		portfolio, err = client.Portfolio(ctx, accounts[accountId-1].ID)
		userPortfolio.Analysis(portfolio)
		mg = "Счёт №" + strconv.Itoa(accountId) + "\n"
	} else if accountId == 0 {
		for _, account := range accounts {
			portfolio, err = client.Portfolio(ctx, account.ID)
			if err != nil {
				fmt.Errorf("Portfolio error: %s", err.Error())
				portfolio, err = client.Portfolio(ctx, account.ID)
				if err != nil {
					mg = "Сейчас брокер не дает полную информацию о портфеле. Такое бывает, попробуйте позже \n"
				}
			}
			userPortfolio.Analysis(portfolio)
		}
	} else {
		return "Счета номер №" + strconv.Itoa(accountId) + " у вас нет"
	}
	userPortfolio.SetPercent()

	return mg + output(&userPortfolio)
}
