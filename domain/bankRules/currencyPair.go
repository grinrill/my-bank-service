package bankRules

import (
	"fmt"
)

// InvalidCurrencyPair returns when it is impossible to do anything
// with source and target currency
type InvalidCurrencyPair struct {source, target Currency}

func (e InvalidCurrencyPair) Error() string {
	return fmt.Sprintf("Currency pair is invalid or has not been implemented yet: %s to %s", e.source, e.target)
}

type currencyPair struct {source, target Currency}

var currencyRates = map[currencyPair]float64{
	{CurrencySBP, CurrencyRUB}: 0.7523,
}
