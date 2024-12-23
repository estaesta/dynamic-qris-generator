// Package qris contains functions to generate and process QRIS.
package qris

import (
	"image"
	// "image/jpeg" for decode image
	_ "image/jpeg"
	// "image/png" for decode image
	_ "image/png"

	"fmt"
	"io"
	"log"
	"strconv"
	"strings"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
)

// ExtractQris extract QRIS data from QRIS image
func ExtractQris(qrisFile io.Reader) (string, error) {
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

// isValid validate QRIS data
func isValid(qrisData string) bool {
	if len(qrisData) < 4 {
		return false
	}

	s := qrisData[len(qrisData)-4:]
	crc := crc(qrisData[:len(qrisData)-4])
	if s != crc {
		return false
	}

	return true
}

func parseQrisData(qrisString string) (Qris, error) {
	// Parse QRIS data
	var qrisData = Qris{}

	// find merchan name between "ID59" and "60"
	idxName := strings.Index(qrisString, "ID59")
	x := qrisString[idxName+4 : idxName+6]
	nName, err := strconv.Atoi(x)
	if err != nil {
		log.Printf("error parsing merchant name length. Err: %v", err)
		return qrisData, err
	}
	qrisData.Name = qrisString[idxName+6 : idxName+6+nName]

	// find merchan city between "60" and "61"
	idxCity := idxName + 6 + nName
	x = qrisString[idxCity+2 : idxCity+4]
	nCity, err := strconv.Atoi(x)
	if err != nil {
		log.Printf("error parsing merchant city length. Err: %v. Got: %v", err, x)
		return qrisData, err
	}
	qrisData.City = qrisString[idxCity+4 : idxCity+4+nCity]

	// find postal code object 61
	idxPostalCode := idxCity + 4 + nCity
	nPostalCode, err := strconv.Atoi(qrisString[idxPostalCode+2 : idxPostalCode+4])
	if err != nil {
		log.Printf("error parsing postal code length. Err: %v", err)
		return qrisData, err
	}
	qrisData.PostalCode = qrisString[idxPostalCode+4 : idxPostalCode+4+nPostalCode]

	return qrisData, nil
}

func qrisEncodeHelper(objNum string, objValue string) string {
	return objNum + fmt.Sprintf("%02d", len(objValue)) + objValue
}

// convertQrisDataToDynamic convert QRIS string static data to dynamic QRIS data
func convertQrisDataToDynamic(qrisString *string, amount int, qrisData ...Qris) error {
	slice := []rune(*qrisString)
	slice[11] = '2'

	// change amount
	*qrisString = string(slice[:12]) + qrisEncodeHelper("54", fmt.Sprintf("%d", amount)) + string(slice[12:])

	if len(qrisData) == 1 {
		x, err := parseQrisData(*qrisString)
		if err != nil {
			return err
		}
		if qrisData[0].Name != "" {
			find := qrisEncodeHelper("59", x.Name)
			repl := qrisEncodeHelper("59", qrisData[0].Name)
			*qrisString = strings.Replace(*qrisString, find, repl, 1)
		}
		if qrisData[0].City != "" {
			find := qrisEncodeHelper("60", x.City)
			repl := qrisEncodeHelper("60", qrisData[0].City)
			*qrisString = strings.Replace(*qrisString, find, repl, 1)
		}
		if qrisData[0].PostalCode != "" {
			find := qrisEncodeHelper("61", x.PostalCode)
			repl := qrisEncodeHelper("61", qrisData[0].PostalCode)
			*qrisString = strings.Replace(*qrisString, find, repl, 1)
		}
	}

	// crc
	slice = []rune(*qrisString)
	*qrisString = string(slice[:len(*qrisString)-4])
	*qrisString += crc(*qrisString)

	return nil
}

// DataToQrisDynamic convert from static QRIS data to dynamic QRIS image
func DataToQrisDynamic(qrisData string, amount int) (image.Image, error) {
	// Validate QRIS data
	if !isValid(qrisData) {
		return nil, fmt.Errorf("invalid QRIS data")
	}

	q, err := parseQrisData(qrisData)
	if err != nil {
		return nil, err
	}

	// Convert QRIS data to dynamic QRIS data
	err = convertQrisDataToDynamic(&qrisData, amount)
	if err != nil {
		return nil, err
	}
	// Generate QRIS image
	img := generateQris(qrisData, q.Name, amount)
	return img, nil
}
