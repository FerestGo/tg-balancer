package balancer

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"

	t "github.com/FerestGo/investapi"
)

var conn *grpc.ClientConn

func InitAnalysis(message string, telegramId int) string {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
	conn, err := grpc.Dial("invest-public-api.tinkoff.ru:443",
		grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{})),
		grpc.WithPerRPCCredentials(oauth.NewOauthAccess(&oauth2.Token{
			AccessToken: token,
		})))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	accounts, err := t.NewUsersServiceClient(conn).GetAccounts(context.Background(), &t.GetAccountsRequest{})

	if err != nil {
		// TODO: обновить все ошибки
		if err.Error() == "can't do request to " {
			return "Сейчас есть проблемы на стороне брокера при получении счетов, попробуйте позже"
		}
		if err.Error() == "bad response to https://api-invest.tinkoff.ru/openapi/user/accounts code=401, body=" {
			return "Неверный токен"
		}
		fmt.Println("Get account error: %s", err)
		return "Сейчас есть не совсем понятная проблема с получением счетов, попробуйте позже с другим токеном." + err.Error()
	}
	var userPortfolio Portfolio
	var portfolio *t.PortfolioResponse

	if accountId != 0 && accounts.Accounts[accountId-1].Id != "" {
		portfolio, err = t.NewOperationsServiceClient(conn).GetPortfolio(ctx, &t.PortfolioRequest{
			AccountId: accounts.Accounts[accountId-1].Id,
		})
		// TODO: обработка err
		userPortfolio.Analysis(portfolio)
		mg = "Счёт №" + strconv.Itoa(accountId) + ": \n\n"
	} else if accountId == 0 {
		for _, account := range accounts.Accounts {
			portfolio, err = t.NewOperationsServiceClient(conn).GetPortfolio(context.Background(), &t.PortfolioRequest{
				AccountId: account.Id,
			})
			// TODO: обработка err
			userPortfolio.Analysis(portfolio)
		}
	} else {
		return "Счета номер №" + strconv.Itoa(accountId) + " у вас нет"
	}
	userPortfolio.SetPercent()

	return mg + Output(&userPortfolio)
}
