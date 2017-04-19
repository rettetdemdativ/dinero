// Author(s): Michael Koeppl

package dinero

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	num float64 = 5000
	cur         = JPY
)

func TestExchangeRate(t *testing.T) {
	var u Currency
	for i := 0; i < len(currencyCodes)-2; i++ {
		u = Currency(i)
		rate, err := u.ExchangeRate(Currency(i + 1))
		if err != nil {
			t.Error(err)
		}
		assert.NotZero(t, rate, "Exchange rate must not be 0")
	}
}

func TestExchangeRateFloat(t *testing.T) {
	// sourceCurrency and targetCurrency defined in amount_test.go
	res, _ := queryAPI(currencyCodes[sourceCurrency], currencyCodes[targetCurrency])
	c := sourceCurrency
	f, _, _ := c.ExchangeRateFloat(targetCurrency)
	assert.Equal(t, res.Rate, f, "Exchange rates in float should match")
}

func TestAmount(t *testing.T) {
	decnum := decimal.NewFromFloat(num)
	c := cur
	a := c.Amount(decnum)
	assert.NotEmpty(t, a, "Created amount must not be empty")
	assert.Equal(t, decimal.NewFromFloat(5000), a.Value, "Amount values should match")
	assert.Equal(t, cur, a.Currency, "Amount currencies should match")
}
