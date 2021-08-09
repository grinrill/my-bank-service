package bankRules

import (
	"math"
	"fmt"
)

// InvalidCurrency returns when it is impossible to do anything
// with source and target currency
type InvalidCurrency struct {currency Currency}

func (e InvalidCurrency) Error() string {
	return fmt.Sprintf("Currency is invalid or has not been implemented yet: %s", e.currency)
}

// Currency Type for currency
type Currency string

// Available currency types
const (
	CurrencySBP Currency = "SBP"
	CurrencyRUB Currency = "RUB"
)

var currencyFractPart = map[Currency]int {
	CurrencyRUB: 2,
	CurrencySBP: 2,
}

// RoundCurrency round currency by defined rules
func RoundCurrency(currency Currency, sum float64) (result float64, err error) {
	fractPart, ok := currencyFractPart[currency]
	if !ok {
		err = InvalidCurrency{currency}
		return
	}

	pow := math.Pow(10, float64(fractPart))

	// Floor because because it is reasonable to round the amount down
	result = math.Floor(sum*pow)/pow
	return
}

// GetCurrencyRate return current currency exchange rate
func GetCurrencyRate(source, target Currency) (rate float64, err error) {
	rate, ok := currencyRates[currencyPair{source, target}]
	if ok {return}

	// If we have no source2target pair
	// Try to get target2source
	rate, ok = currencyRates[currencyPair{target, source}]
	if ok {
		// If ok, convert target2source
		// To source2target rate
		rate = 1/rate
		return
	}

	// Here we have no any pairs with source and target
	err = InvalidCurrencyPair{source, target}
	return
}

// ConvertCurrency will convert source to target currency 
// by current currency exchange rate
func ConvertCurrency(source, target Currency, sum float64) (result float64, err error) {
	rate, err := GetCurrencyRate(source, target)
	if err!=nil {return}

	result = sum*rate
	return
}