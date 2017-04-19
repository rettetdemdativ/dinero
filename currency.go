// Author(s): Michael Koeppl

package dinero

import "github.com/shopspring/decimal"

const (
	// The currencies listed below are the ones supported
	// by the ExchangeRates API and are listed with their
	// symbols and ISO codes.

	// USD -> USDollar ($ / USD)
	USD Currency = iota

	// EUR -> Euro (€ / EUR)
	EUR

	// JPY -> JapaneseYen (¥ / JPY)
	JPY

	// BGN -> BulgarianLev (лв. / BGN)
	BGN

	// CZK -> CzechKoruna (Kč / CZK)
	CZK

	// DKK -> DanishKrone (kr. / DKK)
	DKK

	// GBP -> GreatBritainPound (£ / GBP)
	GBP

	// HUF -> HungarianForint (Ft / HUF)
	HUF

	// LTL -> LithuanianLitas (Lt / LTL)
	LTL

	// LVL -> LatvianLats (Ls / LVL)
	LVL

	// PLN -> PolishZloty (zł / PLN)
	PLN

	// RON -> RomanianLeu (? / RON)
	RON

	// SEK -> SwedishKrona (kr / SEK)
	SEK

	// CHF -> SwissFranc (CHF / CHF)
	CHF

	// NOK -> NorwegianKrone (kr / NOK)
	NOK

	// HRK -> CroatianKuna (kn / HRK)
	HRK

	// RUB -> RussianRubel (₽ / RUB)
	RUB

	// TRY -> TurkishLira (₺ / TRY)
	TRY

	// AUD -> AustralianDollar ($, A$ or AUD / AUD)
	AUD

	// BRL -> BrazilianReal (R$ / BRL)
	BRL

	// CAD -> CanadianDollar ($, Can$, C$ or CAD / CAD)
	CAD

	// CNY -> ChineseYuan (¥ / CNY)
	CNY

	// HKD -> HongKongDollar ($ or HK$ / HKD)
	HKD

	// IDR -> IndonesianRupiah (Rp / IDR)
	IDR

	// ILS -> IsraeliNewShekel (₪  / ILS)
	ILS

	// INR -> IndianRupee (₹ / INR)
	INR

	// KRW -> SouthKoreanWon (₩ / KRW)
	KRW

	// MXN -> MexicanPeso ($ or Mex$ / MXN)
	MXN

	// MYR -> MalaysianRinggit (RM / MYR)
	MYR

	// NZD -> NewZealandDollar ($ / NZD)
	NZD

	// PHP -> PhilippinePeso (₱ / PHP)
	PHP

	// SGD -> SingaporeDollar (S$ or $ / SGD)
	SGD

	// THB -> ThaiBaht (฿ / THB)
	THB

	// ZAR -> SouthAfricanRand (R / ZAR)
	ZAR

	// ISK -> IcelandicKrona (kr or Íkr / ISK)
	ISK
)

var currencyCodes = [...]string{
	"USD",
	"EUR",
	"JPY",
	"BGN",
	"CZK",
	"DKK",
	"GBP",
	"HUF",
	"LTL",
	"LVL",
	"PLN",
	"RON",
	"SEK",
	"CHF",
	"NOK",
	"HRK",
	"RUB",
	"TRY",
	"AUD",
	"BRL",
	"CAD",
	"CNY",
	"HKD",
	"IDR",
	"ILS",
	"INR",
	"KRW",
	"MXN",
	"MYR",
	"NZD",
	"PHP",
	"SGD",
	"THB",
	"ZAR",
	"ISK",
}

var currencySymbols = [...]string{
	"$",
	"€",
	"¥",
	"лв.",
	"Kč",
	"kr.",
	"£",
	"Ft",
	"Lt",
	"Ls",
	"zł",
	"RON",
	"kr",
	"CHF",
	"kr",
	"kn",
	"₽",
	"₺",
	"A$",
	"R$",
	"C$",
	"¥",
	"HK$",
	"Rp",
	"₪",
	"₹",
	"₩",
	"Mex$",
	"RM",
	"$",
	"₱",
	"S$",
	"฿",
	"R",
	"kr",
}

// Currency reprents a currency as an integer value. This is used
// as sort of an enum in which the constant values defined as currencies
// may be used by users of the library and at the same time are used
// as indeces for the currency code array.
type Currency uint8

// ExchangeRate returns the exchance rate when exchanging money from
// the currency c to the target currency or an error.
func (c Currency) ExchangeRate(target Currency) (decimal.Decimal, error) {
	rate, err := rate(currencyCodes[c], currencyCodes[target])
	if err != nil {
		return decimal.Zero, err
	}
	return rate, nil
}

// ExchangeRateFloat returns the exchange rate when exchanging money
// from the currency c to the target currency as a float64 or an error.
func (c Currency) ExchangeRateFloat(target Currency) (rate float64, exact bool, err error) {
	dec, err := c.ExchangeRate(target)
	if err != nil {
		return -1, false, err
	}
	rate, exact = dec.Float64()
	return rate, exact, nil
}

// Amount creates an Amount for the given currency unit and amount.
func (c Currency) Amount(amount decimal.Decimal) Amount {
	return Amount{Value: amount, Currency: c}
}
