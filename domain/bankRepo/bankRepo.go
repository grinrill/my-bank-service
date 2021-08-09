package bankRepo

import (
	"github.com/grinrill/my-bank-service/domain/bankRules"
)

// BankRepo Structure for working with bank accounts database
type BankRepo interface {
	GetAccountCurrency(accountID int) (cur bankRules.Currency, err error)
	GetAccountBalance(accountID int) (balance float64, err error)
	IncrementAccountBalance(accountID int, sum float64) (err error)
	DecrementAccountBalance(accountID int, sum float64) (err error)
}