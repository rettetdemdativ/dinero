// Author(s): Michael Koeppl

package dinero

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	value          float64 = 23
	sourceCurrency         = USD
	targetCurrency         = EUR
)

func TestConvertToMv(t *testing.T) {
	mv := &Amount{decimal.NewFromFloat(value), sourceCurrency}
	resultVal, _ := mv.ConvertToMv(targetCurrency)
	u := sourceCurrency
	rate, _ := u.ExchangeRate(targetCurrency)

	assert.NotEmpty(t, resultVal, "Result must not be empty")
	assert.Equal(t, decimal.NewFromFloat(value).Mul(rate), resultVal.Value, "Results for exchange from USD to EUR should match")
	assert.Equal(t, targetCurrency, resultVal.Currency, "Currencies do not match")
}

func TestConvertTo(t *testing.T) {
	mv := &Amount{decimal.NewFromFloat(value), sourceCurrency}
	amount, err := mv.ConvertTo(targetCurrency)
	t.Log(err)

	u := sourceCurrency
	rate, err := u.ExchangeRate(targetCurrency)
	t.Log(err)
	assert.Equal(t, decimal.NewFromFloat(value).Mul(rate), amount, "Results for exchange from USD to EUR should match")
}

func TestString(t *testing.T) {
	var (
		amount         = 5.254
		expectedResult = "5.254 €"
	)
	a := Amount{Value: decimal.NewFromFloat(amount), Currency: targetCurrency}
	str := a.String()
	assert.Equal(t, expectedResult, str, "Amount strings should match")
}

func TestNewAmount(t *testing.T) {
	a := NewAmount(decimal.NewFromFloat(5000.652), sourceCurrency)
	assert.NotEmpty(t, a, "Created amount must not be empty")
	assert.Equal(t, decimal.NewFromFloat(5000.652), a.Value, "Amount values should match")
	assert.Equal(t, sourceCurrency, a.Currency, "Amount currencies should match")
}

func TestNewAmountFromFloat(t *testing.T) {
	a := NewAmountFromFloat(5000.652, sourceCurrency)
	assert.NotEmpty(t, a, "Created amount must not be empty")
	assert.Equal(t, decimal.NewFromFloat(5000.652), a.Value, "Amount values should match")
	assert.Equal(t, sourceCurrency, a.Currency, "Amount currencies should match")
}

func TestNewAmountFromString(t *testing.T) {
	a, _ := NewAmountFromString("5000.245", sourceCurrency)
	assert.NotEmpty(t, a, "Created amount must not be empty")
	assert.Equal(t, decimal.NewFromFloat(5000.245), a.Value, "Amount values should match")
	assert.Equal(t, sourceCurrency, a.Currency, "Amount currencies should match")
}

func TestNewAmountFromStringError(t *testing.T) {
	_, err := NewAmountFromString("%5000,245", sourceCurrency)
	if err == nil { // error expected
		t.Error("Error expected")
	}
	t.Log(err)
}
