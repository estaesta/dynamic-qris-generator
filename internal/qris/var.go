// Package qris contains functions to generate and process QRIS.
package qris

// Qris struct for QRIS data
type Qris struct {
	Version        string // 00
	Type           string // 01
	AccountInfo    string // 26
	Category       string // 52
	Currency       string // 53
	Amount         string // 54
	CountryCode    string // 58
	Name           string // 59
	City           string // 60
	PostalCode     string // 61
	AdditionalData string // 62
	Crc            string // 63
}

// set constants for data object number
const (
	NumVersion        = "00"
	NumType           = "01"
	NumAccountInfo    = "26"
	NumCategory       = "52"
	NumCurrency       = "53"
	NumAmount         = "54"
	NumCountryCode    = "58"
	NumMerchantName   = "59"
	NumMerchantCity   = "60"
	NumPostalCode     = "61"
	NumAdditionalData = "62"
	NumCrc            = "63"
)
