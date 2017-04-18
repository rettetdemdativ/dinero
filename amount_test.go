// Author(s): Michael Koeppl

package dinero

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	value          float64 = 23
	sourceCurrency         = USD
	targetCurrency         = EUR
)

func TestConvertToMv(t *testing.T) {
	mv := &Amount{value, sourceCurrency}
	resultVal, _ := mv.ConvertToMv(targetCurrency)
	u := sourceCurrency
	rate, _ := u.ExchangeRate(targetCurrency)

	assert.NotEmpty(t, resultVal, "Result must not be empty")
	assert.Equal(t, value*rate, resultVal.Value, "Results for exchange from USD to EUR should match")
	assert.Equal(t, targetCurrency, resultVal.Currency, "Currencies do not match")
}

func TestConvertTo(t *testing.T) {
	mv := &Amount{value, sourceCurrency}
	amount, err := mv.ConvertTo(targetCurrency)
	t.Log(err)

	u := sourceCurrency
	rate, err := u.ExchangeRate(targetCurrency)
	t.Log(err)
	assert.Equal(t, value*rate, amount, "Results for exchange from USD to EUR should match")
}

func TestString(t *testing.T) {
	var (
		amount         float64 = 5
		expectedResult         = fmt.Sprintf("%f €", amount)
	)
	a := Amount{Value: amount, Currency: targetCurrency}
	str := a.String()
	assert.Equal(t, expectedResult, str, "Amount strings should match")
}
