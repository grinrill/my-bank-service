package bankService

import (
	"database/sql"

	"github.com/grinrill/my-bank-service/domain/bankService"
	"github.com/grinrill/my-bank-service/domain/bankRules"
	"github.com/grinrill/my-bank-service/domain/bankRepo"
)

// Account implements bankService.AccountInterface
type Account struct {
	accountID int
	bankRepo  bankRepo.BankRepo
}

// AddFunds implements bankService.AccountInterface
func (a *Account) AddFunds(sum float64) (err error) {
	err = a.bankRepo.IncrementAccountBalance(a.accountID, sum)
	if err!=nil {return}

	return a.SumProfit()
}

// SumProfit implements bankService.AccountInterface
func (a *Account) SumProfit() (err error) {
	sumProfitRate := bankRules.GetSumProfitRate()

	balance, err := a.bankRepo.GetAccountBalance(a.accountID)
	if err!=nil {return}
	currency, err := a.bankRepo.GetAccountCurrency(a.accountID)
	if err!=nil {return}

	sum, err := bankRules.RoundCurrency(currency, balance*sumProfitRate)
	if err!=nil {return}

	err = a.bankRepo.IncrementAccountBalance(accountID, increment)
	return
}