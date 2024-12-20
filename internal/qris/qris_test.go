// Description: This file is used to generate process QRIS and generate QRIS image.
package qris

import (
	"testing"

	_ "github.com/makiuchi-d/gozxing"
	_ "github.com/makiuchi-d/gozxing/qrcode"
)

func Test_parseQris(T *testing.T) {
	x, _ := parseQrisData("WWW0215ID10222210372890303UKE5204504553033605802ID5905Twnku6006Bekasi610517610630463A9")
	if x.Name != "Twnku" {
		T.Errorf("Expected Twnku, got %s", x.Name)
	}
	if x.City != "Bekasi" {
		T.Errorf("Expected Bekasi, got %s", x.City)
	}
	if x.PostalCode != "17610" {
		T.Errorf("Expected 11111, got %s", x.PostalCode)
	}
}
