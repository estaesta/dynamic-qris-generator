// Package qris contains functions to generate and process QRIS.
package qris

import (
	"image"
	"image/color"
	"image/draw"
	"log"
	"strings"

	// dependency for jpeg encoding required by gozxing
	_ "image/jpeg"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gobold"
	"golang.org/x/image/math/fixed"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// generateQris generate a QRIS image from string data
func generateQris(qrisData, name string, amount int) image.Image {
	// Generate QRIS data
	enc := qrcode.NewQRCodeWriter()
	qr, err := enc.Encode(qrisData, gozxing.BarcodeFormat_QR_CODE, 500, 500, nil)
	if err != nil {
		log.Fatalf("error encoding QRIS data. Err: %v", err)
	}

	// Create base image
	img := image.NewRGBA(image.Rect(0, 0, 500, 650))
	draw.Draw(img, img.Bounds(), image.White, image.Point{}, draw.Src)
	// Place QRIS image on base image
	draw.Draw(img, img.Bounds(), qr, image.Point{0, -100}, draw.Over)
	// Add label to QRIS image
	addLabel(img, name, amount)

	return img
}

// addLabel add label to QRIS image including name and amount
func addLabel(img *image.RGBA, name string, amount int) {
	stringsToPrint := []string{}
	name = "a.n " + strings.ToUpper(name)
	col := color.Black

	ft, err := freetype.ParseFont(gobold.TTF)
	if err != nil {
		log.Fatalf("error parsing font. Err: %v", err)
	}

	d := &font.Drawer{
		Dst:  img,
		Src:  image.NewUniform(col),
		Face: truetype.NewFace(ft, &truetype.Options{Size: 24}),
	}

	dst := d.MeasureString(name)

	if dst.Ceil() > img.Bounds().Dx()-100 {
		maxWidth := img.Bounds().Dx() - 100
		stringsToPrint = splitByRenderedWidth(name, fixed.I(maxWidth), d)
	} else {
		stringsToPrint = append(stringsToPrint, name)
	}

	// Measure total height of all strings
	totalHeight := len(stringsToPrint) * 30 // 30 is the line height
	startY := (200 - totalHeight) / 2       // Center Y coordinate

	// Limiting the number of lines to 3
	if len(stringsToPrint) > 3 {
		lastLine := stringsToPrint[len(stringsToPrint)-1]
		lastLineWords := strings.Split(lastLine, " ")
		lastWord := lastLineWords[len(lastLineWords)-1]
		stringsToPrint = stringsToPrint[:2]
		stringsToPrint = append(stringsToPrint, "... "+lastWord)
	}

	// Draw the strings
	for i, str := range stringsToPrint {
		dst = d.MeasureString(str)
		d.Dot = fixed.Point26_6{
			X: fixed.I(img.Bounds().Dx()/2) - dst/2, // Center horizontally
			Y: fixed.I(startY + i*30),               // Center vertically
		}
		d.DrawString(str)
	}

	// Draw the amount
	p := message.NewPrinter(language.Indonesian)
	dst = d.MeasureString(p.Sprintf("Rp %d", amount))
	d.Dot = fixed.Point26_6{
		X: fixed.I(img.Bounds().Dx()/2) - dst/2,
		Y: fixed.I(img.Bounds().Dy() - 50),
	}
	d.DrawString(p.Sprintf("Rp %d", amount))
}

func splitByRenderedWidth(input string, maxWidth fixed.Int26_6, drawer *font.Drawer) []string {
	var result []string
	var currentLine string
	var currentWidth fixed.Int26_6

	words := strings.Split(input, " ")

	for _, word := range words {
		wordWidth := drawer.MeasureString(word + " ")

		if currentWidth+wordWidth > maxWidth && len(currentLine) > 0 {
			// Finalize the current line
			result = append(result, currentLine)
			currentLine = ""
			currentWidth = 0
		}

		if currentLine != "" {
			currentLine += " "
		}
		currentLine += word
		currentWidth += wordWidth
	}

	// Append the last line
	if len(currentLine) > 0 {
		result = append(result, currentLine)
	}

	return result
}
