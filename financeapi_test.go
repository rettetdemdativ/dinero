// Author(s): Michael Koeppl

package dinero

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	sourceCurrencyStr = "USD"
	targetCurrencyStr = "EUR"
)

func TestMakeUrl(t *testing.T) {
	url := URL(sourceCurrencyStr, targetCurrencyStr)
	expectedResult := "http://rate-exchange-1.appspot.com/currency?from=USD&to=EUR"
	assert.Equal(t, url, expectedResult, "Generated URL should match expected URL")
}

func TestCustomClient(t *testing.T) {
	cc := customClient()
	assert.Equal(t, cc.Timeout, timeout, "Set timeout of the custom client should be 5 seconds")
}

func TestQueryAPI(t *testing.T) {
	res, _ := queryAPI(sourceCurrencyStr, targetCurrencyStr)

	assert.NotEmpty(t, res, "Response must not be empty")

	assert.NotZero(t, res.Rate, "Result rate must not be 0")
}

func TestRate(t *testing.T) {
	r, _ := rate(sourceCurrencyStr, targetCurrencyStr)
	assert.NotZero(t, r, "Exchange rate from response must not be 0")
}
