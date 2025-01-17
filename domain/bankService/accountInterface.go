package bankService

import (
	"github.com/grinrill/my-bank-service/domain/bankRules"
)

// AccountInterface Bank account interface
type AccountInterface interface {
	// AddFunds Позволяет внести на счёт сумму sum
	AddFunds(sum float64) error
	// SumProfit Рассчитывает процент по вкладу и полученные деньги вносит на счёт
	SumProfit() error
	// Withdraw Производит списание со счёта по указанным правилам. Если списание выходит за рамки правил, выдаёт ошибку
	Withdraw(sum float64) error
	// GetCurrency Выдаёт валюту счёта
	GetCurrency() (bankRules.Currency, error)
	// GetAccountCurrencyRate Выдаёт курс валюты счёта к передаваемой валюте cur
	GetAccountCurrencyRate(cur bankRules.Currency) (float64, error)
	// GetBalance Выдаёт баланс счёта в указанной валюте
	GetBalance(cur string) (float64, error)
}
