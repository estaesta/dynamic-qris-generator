// Description: This file is used to generate process QRIS and generate QRIS image.
package qris

import (
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"

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
