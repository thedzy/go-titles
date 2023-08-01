package main

import (
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/golang/freetype"
	// "golang.org/x/image/font/gofont/goitalic"
	"golang.org/x/image/font/opentype"
	// "github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	// "golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/math/fixed"
	"golang.org/x/sys/unix"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"strings"
)

var (
	text                    = flag.String("text", "Hello World!", "text to render")
	displayCharactersOption = flag.String("characters", " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}", "text to render")
	displayResolutionOption = flag.Int("resolution", 20, "text to render")
	displayFontOption       = flag.String("render-font", "/System/Library/Fonts/Monaco.ttf", "filename of the ttf font")
	pixelAspect             = flag.Float64("aspect", 0.66, "character height to width")
	fontName                = flag.String("font", "/System/Library/Fonts/Supplemental/Arial.ttf", "filename of the ttf font")
	fontSize                = flag.Float64("size", 300.0, "font size in points")
	loadFile                = flag.String("load", "", "load saved character map")
	saveFile                = flag.String("save", "", "save character map")

	useInverted = flag.Bool("allow-inverted", false, "use inverted characters")
	debug       = flag.Bool("debug", false, "debug mode")
)

func main() {
	flag.Parse()

	displayResolution := *displayResolutionOption
	displayCharacters := *displayCharactersOption
	displayFont := *displayFontOption

	// Load character map
	var characterMap map[string][][]int
	if *loadFile != "" {
		if *debug {
			fmt.Println("Loading File")
		}
		if strings.HasSuffix(*loadFile, ".json") {
			jsonData, err := ioutil.ReadFile(*loadFile)
			if err != nil {
				log.Fatal(err)
			}
			err = json.Unmarshal(jsonData, &characterMap)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			characterMap, _ = loadCharacterMapFromDisk(*loadFile)
		}
		// Get resolution from loaded map
		for _, value := range characterMap {
			displayResolution = len(value)
			break
		}
	} else {
		if *debug {
			fmt.Printf(displayCharacters)
		}
	}
	*fontSize = *fontSize * float64(displayResolution)

	// Draw title
	renderImage, textImageRect := renderText(*text, *fontSize, 72.0, *fontName, image.White, image.Black)
	croppedImage := cropImageToDimension(renderImage, textImageRect.X, textImageRect.Y)

	// Check if the image width is greater than window width
	winSize, err := getWinSize()
	if err != nil {
		fmt.Println(err)
		winSize.Col = 1200
	}
	displayWidth := int(winSize.Col) * displayResolution
	if int(textImageRect.X>>6) > displayWidth {
		if *debug {
			fmt.Println("Scaling")
		}
		// Calculate the scale factor to resize the image to width 120
		scale := float64(displayWidth) / float64(textImageRect.X>>6)
		croppedImage = scaleImageToDimension(croppedImage, int(displayWidth), int(float64(textImageRect.Y>>6)*scale))
	}
	if *debug {
		saveImage(croppedImage, "test.png")
	}

	// Scale to pixel ratio
	bounds := croppedImage.Bounds()
	croppedImage = scaleImageToDimension(croppedImage, bounds.Max.X, int(float64(bounds.Max.Y)**pixelAspect))

	// Build and fill a 4D matrix of the title
	brightnessMatrix := getEmptyBrightnessMatrix(croppedImage, displayResolution)
	brightnessMatrix = fillBrightnessMatrix(brightnessMatrix, croppedImage, displayResolution)

	// Build character map if not loading from file
	if *loadFile == "" {
		characterMap = mapCharacters(displayCharacters, displayFont, displayResolution)
	}

	// Save character map
	if *saveFile != "" {
		if strings.HasSuffix(*saveFile, ".json") {
			fmt.Printf("Save json to disk")

			// Open a file for writing
			file, err := os.Create(*saveFile)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			// Create a JSON encoder to write to the file
			encoder := json.NewEncoder(file)

			// Encode the characterMap and write it to the file
			err = encoder.Encode(characterMap)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Save map to disk")
			saveCharacterMapToDisk(characterMap, *saveFile)
		}
	}

	// Debug characters
	if *debug {
		for key, value := range characterMap {
			fmt.Printf("Key: %s\n", key)
			drawMatrix(value)
		}
		os.Exit(1)
	}

	// Render to screen
	rows, cols := len(brightnessMatrix), len(brightnessMatrix[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			character, invert := findMatch(characterMap, brightnessMatrix[x][y])
			if invert {
				fmt.Print("\x1B[7m")
			}
			fmt.Print(character)

			fmt.Print("\x1B[0m")
		}
		fmt.Println()
	}
}

// findMatch Find a character that matches a section of image
func findMatch(characterMap map[string][][]int, testMatrix [][]int) (string, bool) {
	match := " "
	bestScore := 255.0
	invertedMatch := " "
	invertedBestScore := 0.0
	inverted := false
	for character, matrix := range characterMap {
		score := calculateMSE(testMatrix, matrix)
		if score < bestScore {
			bestScore = score
			match = character
		}
		if score > invertedBestScore {
			invertedBestScore = score
			invertedMatch = character
		}
	}

	if *useInverted {
		if bestScore <= 255-invertedBestScore {
			// fmt.Printf("%0.0f v %0.0f, regular\n", bestScore, 255-invertedBestScore)
			inverted = false
		} else {
			// fmt.Printf("%0.0f v %0.0f, inverted\n", bestScore, 255-invertedBestScore)
			match = invertedMatch
			inverted = true
		}
	}

	return match, inverted
}

// RenderText to image
func renderText(text string, fontSize float64, imageDPI float64, fontName string, foreground *image.Uniform, background *image.Uniform) (image.Image, fixed.Point26_6) {
	// Initialize the context.
	render := image.NewRGBA(image.Rect(0, 0, int(fontSize)*len(text), int(fontSize*1.2)))
	draw.Draw(render, render.Bounds(), background, image.ZP, draw.Src)

	var textImageRect fixed.Point26_6
	fontBytes, err := ioutil.ReadFile(fontName)
	if err != nil {
		log.Println(err)
		return nil, fixed.P(0, 0)
	}
	if strings.HasSuffix(fontName, ".ttf2") {

		fontData, err := freetype.ParseFont(fontBytes)
		if err != nil {
			log.Println(err)
			return nil, fixed.P(0, 0)
		}
		newContext := freetype.NewContext()
		newContext.SetDPI(imageDPI)
		newContext.SetFont(fontData)
		newContext.SetFontSize(fontSize)
		newContext.SetClip(render.Bounds())
		newContext.SetDst(render)
		newContext.SetSrc(foreground)
		newContext.SetHinting(font.HintingFull)

		// Draw the text.
		text = strings.Replace(fontName, "\t", "  ", -1) // convert tabs into spaces
		pt := freetype.Pt(0, 0+int(newContext.PointToFixed(fontSize)>>6))

		textImageRect, err = newContext.DrawString(text, pt)
		if err != nil {
			log.Println(err)
			return nil, fixed.P(0, 0)
		}
		pt.Y += newContext.PointToFixed(fontSize)
	} else {
		fontData, err := opentype.Parse(fontBytes)
		if err != nil {
			log.Fatalf("Parse: %v", err)
		}
		face, err := opentype.NewFace(fontData, &opentype.FaceOptions{
			Size:    fontSize,
			DPI:     imageDPI,
			Hinting: font.HintingNone,
		})
		if err != nil {
			log.Fatalf("NewFace: %v", err)
		}

		fontDrawer := font.Drawer{
			Dst:  render,
			Src:  foreground,
			Face: face,
			Dot:  fixed.P(0, int(fontSize)),
		}

		fontDrawer.DrawString(text)
		// Determine the rendering bounds of the text
		bounds, _ := fontDrawer.BoundString(text)
		textImageRect = fixed.P(int(bounds.Min.X/64), int(bounds.Max.Y/64))
	}

	return render, textImageRect
}

// cropImageToDimension Crops image to the dimensions
func cropImageToDimension(img image.Image, x, y fixed.Int26_6) image.Image {
	// Create a rectangle using the fixed-point dimensions
	rect := image.Rect(0, 0, x.Floor(), int(float32(y.Floor())*1.5))

	// Create a new RGBA image with the specified dimensions
	cropped := image.NewRGBA(rect)

	// Copy the cropped region from the original image to the new image
	draw.Draw(cropped, cropped.Bounds(), img, rect.Min, draw.Src)

	return cropped
}

// scaleImageToDimension Scales image to the dimensions
func scaleImageToDimension(img image.Image, x, y int) image.Image {
	scaled := image.NewRGBA(image.Rect(0, 0, int(x), int(y)))
	draw.NearestNeighbor.Scale(scaled, scaled.Rect, img, img.Bounds(), draw.Over, nil)
	return scaled
}

// scaleImageToProportions scales the image to the size while maintaining proportions
func scaleImageToProportions(img image.Image, size int, method int) image.Image {
	bounds := img.Bounds()
	x, y := bounds.Max.X, bounds.Max.Y

	var newSize int
	if x > y {
		newSize = x
	} else {
		newSize = y
	}
	// Create a rectangle using the fixed-point dimensions
	rect := image.Rect(0, 0, newSize, newSize)

	// Create a new RGBA image with the specified dimensions
	resized := image.NewRGBA(rect)

	// Calculate offsets for centering the resized image
	offsetX := -(newSize - img.Bounds().Dx()) / 2
	offsetY := -(newSize - img.Bounds().Dy()) / 2
	draw.Draw(resized, resized.Bounds(), img, image.Point{offsetX, offsetY}, draw.Src)

	scaled := image.NewRGBA(image.Rect(0, 0, size, size))

	switch method {
	case 1:
		draw.ApproxBiLinear.Scale(scaled, scaled.Rect, resized, resized.Bounds(), draw.Over, nil)
	case 2:
		draw.NearestNeighbor.Scale(scaled, scaled.Rect, resized, resized.Bounds(), draw.Over, nil)
	case 3:
		draw.CatmullRom.Scale(scaled, scaled.Rect, resized, resized.Bounds(), draw.Over, nil)
	default:
		draw.BiLinear.Scale(scaled, scaled.Rect, resized, resized.Bounds(), draw.Over, nil)
	}
	return scaled
}

// getEmptyBrightnessMatrix reads an image file and returns an empty 4D matrix.
func getEmptyBrightnessMatrix(img image.Image, resolution int) [][][][]int {

	// Get the image dimensions
	bounds := img.Bounds()

	// Define the dimensions of the 2D matrix
	rows, cols := bounds.Max.Y/resolution, bounds.Max.X/resolution

	// Create the 2D matrix of resXres matrices
	matrix := make([][][][]int, rows)

	// Initialize each element of the  matrix
	for x := 0; x < rows; x++ {
		matrix[x] = make([][][]int, cols)
		for y := 0; y < cols; y++ {
			matrix[x][y] = make([][]int, resolution)
			for r := 0; r < resolution; r++ {
				matrix[x][y][r] = make([]int, resolution)
			}
		}
	}

	return matrix
}

// fillBrightnessMatrix reads an image file and fills a matrix of brightness values.
func fillBrightnessMatrix(matrix [][][][]int, img image.Image, resolution int) [][][][]int {

	// Get the image dimensions
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	// Define the dimensions of the 2D matrix
	rows, cols := height/resolution, width/resolution

	// Initialize each element as a resolution x resolution matrix
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			for x1 := 0; x1 != resolution; x1++ {
				for y1 := 0; y1 != resolution; y1++ {
					// Get the color of the pixel
					rgba := img.At((y*resolution)+y1, (x*resolution)+x1).(color.RGBA)

					// Calculate the brightness value
					brightness := (uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3

					// Store the brightness value in the matrix
					matrix[x][y][x1][y1] = int(brightness)
				}
			}
			if *debug {
				drawMatrix(matrix[x][y])
			}
		}
	}

	return matrix
}

// mapCharacters Create a brightness map of the chracters
func mapCharacters(characters string, fontName string, resolution int) map[string][][]int {
	characterMap := make(map[string][][]int)

	for _, character := range characters {
		renderCharacter, dimensions := renderText(string(character), 20, 72, fontName, image.Black, image.Transparent)
		croppedCharacter := scaleImageToDimension(renderCharacter, int(dimensions.Y>>6), int(dimensions.Y>>6))
		croppedCharacter = scaleImageToProportions(renderCharacter, resolution, 2)
		if *debug {
			saveImage(croppedCharacter, fmt.Sprintf("%d.png", character))
		}

		characterMap[string(character)] = getImageMatrix(croppedCharacter)
	}

	return characterMap
}

// getImageMatrix Get the matrix of brightness values of an image
func getImageMatrix(img image.Image) [][]int {
	bounds := img.Bounds()
	resolution := bounds.Max.X

	// Build a matrix
	characterMatrix := make([][]int, resolution)
	for i := 0; i < resolution; i++ {
		characterMatrix[i] = make([]int, resolution)
	}

	// Fill bright values
	for x := 0; x < resolution; x++ {
		for y := 0; y < resolution; y++ {
			rgba := img.At(y, x).(color.RGBA)
			// Calculate the brightness value
			// brightness := (uint32(rgba.R) + uint32(rgba.G) + uint32(rgba.B)) / 3
			brightness := uint32(rgba.A)

			// Store the brightness value in the matrix
			characterMatrix[x][y] = int(brightness)
		}
	}

	return characterMatrix
}

func calculateMSE(matrix1, matrix2 [][]int) float64 {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		fmt.Printf("%d,%d != %d,%d\n", len(matrix1), len(matrix1[0]), len(matrix2), len(matrix2[0]))
		fmt.Println(matrix2)
		panic("Matrices must have the same dimensions.")
	}

	rows, cols := len(matrix1), len(matrix1[0])
	sum := 0.0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			diff := matrix1[i][j] - matrix2[i][j]
			sum += float64(diff * diff)
		}
	}

	mse := sum / float64(rows*cols)
	return math.Sqrt(mse) // The score ranges from 0 (match) to 255 (exact opposite)
}

// saveCharacterMapToDisk Saves the map to disk to save render time and distribute styles
func saveCharacterMapToDisk(data map[string][][]int, filename string) error {
	// Create or open the file for writing
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a new encoder for writing
	encoder := gob.NewEncoder(file)

	// Encode and save the map to the file
	err = encoder.Encode(data)
	if err != nil {
		return err
	}

	fmt.Println("Character map saved to file:", filename)
	return nil
}

// loadCharacterMapFromDisk Loads the map from disk of a saved render
func loadCharacterMapFromDisk(filename string) (map[string][][]int, error) {
	// Open the file for reading
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a new decoder for reading
	decoder := gob.NewDecoder(file)

	// Decode and load the map from the file
	var data map[string][][]int
	err = decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// getWinSize Get the full size of the window (*nix/mac)
func getWinSize() (*unix.Winsize, error) {
	// Get window dimensions for Unix
	winDimensions, err := unix.IoctlGetWinsize(int(os.Stdout.Fd()), unix.TIOCGWINSZ)
	if err != nil {
		return nil, err
	}
	return winDimensions, nil
}

// saveImage Debugging, save image to disk
func saveImage(img image.Image, name string) {
	output, _ := os.Create(name)
	defer output.Close()

	png.Encode(output, img)
	fmt.Printf("Image saved to %s\n", name)
}

// drawMatrix Debugging, draw a matrix
func drawMatrix(value [][]int) {
	resolution := len(value)
	for r := 0; r < resolution; r++ {
		var formatted []string
		for _, num := range value[r] {
			formatted = append(formatted, fmt.Sprintf("%3d", num))
		}
		fmt.Println("[" + strings.Join(formatted, " ") + "]")
	}
	for y := 0; y < resolution; y++ {
		for x := 0; x < resolution; x++ {
			matrixValue := value[y][x]
			if matrixValue > 192 {
				fmt.Print("X")
			} else if matrixValue > 128 {
				fmt.Print("x")
			} else if matrixValue > 64 {
				fmt.Print(".")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}