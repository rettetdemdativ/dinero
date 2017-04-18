// Author(s): Michael Koeppl

package dinero

import (
	"fmt"
)

// Amount represents a money value in a given currency.
type Amount struct {
	Value    float64
	Currency Currency
}

// ConvertToMv converts mv to another target currency. It returns the converted
// money value or an error.
func (a Amount) ConvertToMv(target Currency) (*Amount, error) {
	resval := &Amount{}
	rate, err := rate(currencyCodes[a.Currency], currencyCodes[target])
	if err != nil {
		return nil, err
	}
	resval.Value = a.Value * rate
	resval.Currency = target
	return resval, nil
}

// ConvertTo converts mv to another target currency. It returns the converted
// money value as a float64 or an error.
func (a Amount) ConvertTo(target Currency) (float64, error) {
	res, err := a.ConvertToMv(target)
	if err != nil {
		return -1, err
	}
	return res.Value, nil
}

// String returns the amount with its currency symbol in the
// predefined format "<value> <symbol>".
func (a Amount) String() string {
	return fmt.Sprintf("%f %s", a.Value, currencySymbols[a.Currency])
}
