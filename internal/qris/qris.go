// Description: This file is used to generate process QRIS and generate QRIS image.
package qris

import (
	"io"

	_ "github.com/makiuchi-d/gozxing"
	_ "github.com/makiuchi-d/gozxing/qrcode"
)

// A function to read QRIS from QRIS image to string
func ReadQris(qrisFile io.Reader) string {
	var qrisData string
	// Read QRIS data
	// img, _, _ := image.Decode(qrisFile)
	return qrisData
}

func validateQrisData(qrisData string) bool {
	// Validate QRIS data
	return true
}

// A function to convert QRIS data to dynamic QRIS data
func convertQrisDataToDynamic(qrisData *string) {
	// Convert QRIS data
}

// A function to generate QRIS image from QRIS data string
func generateQris(qrisData string) []byte {
	// Generate QRIS data
	return nil
}

// A function to convert from static QRIS image to dynamic QRIS image
func QrisToQrisDynamic(qrisImage []byte) ([]byte, error) {
	return nil, nil
}

// A function to convert from static QRIS data to dynamic QRIS image
func DataToQrisDynamic(qrisData string) ([]byte, error) {
	return nil, nil
}
