package util

// Constants for all supported currencies
const (
	USD = "USD"
	EUR = "EUR"
	JPY = "JPY"
)

// IsSupportedCurrency returns true if the currency is supported
func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, JPY:
		return true
	}
	return false
}
