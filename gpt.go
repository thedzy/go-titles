package main

import (
	"fmt"
	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
	"image"
	"image/png"
	"os"
)

func main() {
	// Get the font data from the gomono package.
	fontData, err := truetype.Parse(gomono.TTF)
	if err != nil {
		fmt.Println("Error parsing font data:", err)
		return
	}

	// Create a new image to draw on
	img := image.NewRGBA(image.Rect(0, 0, 300, 100)) // Increase the width to make sure text fits

	// Create a new freetype context and set the font and font size
	f := fixed.Int26_6(32 * 64) // font size
	c := freetype.NewContext()
	c.SetFont(fontData)
	c.SetFontSize(float64(f) / 64) // Convert fixed.Int26_6 to float64

	// Set the destination image and foreground color
	c.SetDst(img)
	c.SetSrc(image.Black)

	// Center the text horizontally and vertically within the image bounds
	text := "Hello, Golang!"
	pt := freetype.Pt((img.Bounds().Dx()-c.PointToFixed(float64(len(text))).Ceil())/2, img.Bounds().Dy()/2)

	// Draw the text on the image
	_, err = c.DrawString(text, pt)
	if err != nil {
		fmt.Println("Error drawing text:", err)
		return
	}

	// Save the image as PNG
	file, err := os.Create("output.png")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		fmt.Println("Error encoding image:", err)
		return
	}
}
