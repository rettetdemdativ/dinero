// Author(s): Michael Koeppl

package dinero

import "testing"
import "github.com/stretchr/testify/assert"

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

func TestAmount(t *testing.T) {
	const (
		num float64 = 50000
		cur         = JPY
	)
	c := cur
	a := c.Amount(num)
	assert.NotEmpty(t, a, "Created amount must not be empty")
	assert.Equal(t, num, a.Value, "Amount values should match")
	assert.Equal(t, cur, a.Currency, "Amount currencies should match")
}
