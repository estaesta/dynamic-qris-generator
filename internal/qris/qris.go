// Description: This file is used to generate process QRIS and generate QRIS image.
package qris

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"

	"github.com/makiuchi-d/gozxing"
	_ "github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	_ "github.com/makiuchi-d/gozxing/qrcode"
)

// A function to read QRIS from QRIS image to string
func ReadQris(qrisFile io.Reader) (string, error) {
	// Read QRIS data
	img, _, err := image.Decode(qrisFile)
	if err != nil {
		return "", err
	}
	// prepare BinaryBitmap
	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return "", err
	}

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return "", err
	}

	return result.GetText(), nil
}

func validateQrisData(qrisData string) bool {
	// Validate QRIS data
	return true
}

type QrisData struct {
	Version                    string // 00
	Type                       string // 01
	MerchantAccountInformation string // 26
	MerchantCategoryCode       string // 52
	TransactionCurrency        string // 53
	CountryCode                string // 58
	MerchantName               string // 59
	MerchantCity               string // 60
	PostalCode                 string // 61
	MerchantAdditionalData     string // 62
	Crc                        string // 63
}

func parseQrisData(qrisString string) QrisData {
	// Parse QRIS data
	var qrisData = QrisData{}

	// find between "ID59" and "60"
	qrisData.MerchantName = qrisString[strings.Index(qrisString, "ID59")+4 : strings.Index(qrisString, "60")]

	return qrisData
}

// A function to convert QRIS data to dynamic QRIS data
func convertQrisDataToDynamic(qrisData *string) error {
	// Convert QRIS data
	return nil
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
