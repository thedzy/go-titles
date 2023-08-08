package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/image/font/gofont/gomono"

	// "github.com/golang/freetype"
	// "github.com/golang/freetype/truetype"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/math/fixed"
	"golang.org/x/sys/unix"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"strings"
)

var (
	// Render options
	text              = flag.String("text", "Hello World!", "text to render")
	displayCharacters = flag.String("characters", " !\"#$%&'()*+,-./0123456789:;<=>?@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}", "text to render, ignored when loading a map")
	displayResolution = flag.Int("resolution", 16, "text to render, ignored when loading a map")
	pixelAspect       = flag.Float64("aspect", 0.5, "character height to width")
	fontName          = flag.String("font", "/System/Library/Fonts/Supplemental/Arial.ttf", "filename of the ttf font")
	fontSize          = flag.Float64("size", 25.0, "font size in points")
	maxWidth          = flag.Int("max-width", 0, "maximium width to render")
	useInverted       = flag.Bool("allow-inverted", false, "use inverted characters, ignored when writing to file")
	renderMode        = flag.Int("mode", 20, "render mode")

	// File options
	loadFile   = flag.String("load", "", "load saved character map")
	saveFile   = flag.String("save", "", "save character map")
	outputFile = flag.String("output", "", "save output to file")

	debug = flag.Bool("debug", false, "debug mode")
)

func main() {
	flag.Parse()
	displayCharacters = removeSpecialCharactersAndDuplicates(*displayCharacters)

	// Set output
	screenOutput := true
	if *outputFile != "" && isValidSavePath(*outputFile) {
		screenOutput = false
		*useInverted = false
	} else if !isValidSavePath(*outputFile) {
		fmt.Printf("%s, is not a valid path\n", *outputFile)
		os.Exit(1)
	}

	// Load character map
	var characterMap map[string][][]int
	if *loadFile != "" && isValidFilePath(*loadFile) {
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
			*displayResolution = len(value)
			break
		}
	} else if *loadFile == "" {
		if *debug {
			fmt.Println(*displayCharacters)
		}
	} else if *loadFile != "" && !isValidFilePath(*loadFile) {
		fmt.Printf("%s, is not a valid path\n", *loadFile)
		os.Exit(1)
	}
	*fontSize = *fontSize * float64(*displayResolution)

	// Draw title
	fontBytes := getFontBytes(*fontName)
	_ = fontBytes
	renderImage, textImageRect := renderText(*text, fontBytes, *fontSize, 72.0, image.White, image.Transparent)
	croppedImage := cropImageToDimension(renderImage, 0, 0, int(textImageRect.X>>6), int(textImageRect.Y>>6))

	// Check if the image width is greater than window width and if we are rending to screen
	if *maxWidth == 0 {
		if screenOutput {
			winSize, err := getWinSize()
			if err != nil {
				fmt.Println(err)
				*maxWidth = math.MaxInt32
			}
			*maxWidth = int(winSize.Col)
		} else {
			*maxWidth = math.MaxInt32
		}
	}
	displayWidth := int(*maxWidth) * *displayResolution
	if int(textImageRect.X>>6) > displayWidth {
		if *debug {
			fmt.Println("Scaling")
		}
		// Calculate the scale factor to resize the image to width 120
		scale := float64(displayWidth) / float64(textImageRect.X>>6)
		croppedImage = scaleImageToDimension(croppedImage, int(displayWidth), int(float64(textImageRect.Y>>6)*scale), 0)
	}
	if *debug {
		saveImage(croppedImage, "title.png")
	}

	// Scale to pixel ratio
	bounds := croppedImage.Bounds()
	croppedImage = scaleImageToDimension(croppedImage, bounds.Max.X, int(float64(bounds.Max.Y)**pixelAspect), 0)
	croppedImage = cropBlank(croppedImage)

	// Build and fill a 4D matrix of the title
	brightnessMatrix := getEmptyBrightnessMatrix(croppedImage, *displayResolution)
	brightnessMatrix = fillBrightnessMatrix(brightnessMatrix, croppedImage, *displayResolution)

	// Build character map if not loading from file
	if !isValidFilePath(*loadFile) {
		characterMap = mapCharacters(*displayCharacters, *displayResolution)
	}

	// Save character map
	if *saveFile != "" && isValidSavePath(*saveFile) {
		if strings.HasSuffix(*saveFile, ".json") {
			fmt.Println("Save json to disk")

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
	} else if !isValidSavePath(*saveFile) {
		fmt.Printf("%s, is not a valid path\n", *saveFile)
		os.Exit(2)
	}

	// Debug characters
	if *debug {
		for key, value := range characterMap {
			fmt.Printf("Key: %s %d\n", key, []rune(key)[0])
			drawMatrix(value)
		}
		os.Exit(255)
	}

	// Render to screen/file
	rows, cols := len(brightnessMatrix), len(brightnessMatrix[0])
	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			character, invert := findMatch(characterMap, brightnessMatrix[x][y], *renderMode)
			if invert {
				fmt.Print("\x1B[7m")
			}
			if screenOutput {
				fmt.Print(character)
			} else {
				outputToFile(*outputFile, character)
			}
			fmt.Print("\x1B[0m")
		}
		if screenOutput {
			fmt.Println()
		} else {
			outputToFile(*outputFile, "\n")
		}
	}
}

// findMatch Find a character that matches a section of image
func findMatch(characterMap map[string][][]int, testMatrix [][]int, method int) (string, bool) {
	match := " "
	bestScore := 255.0
	invertedMatch := " "
	invertedBestScore := 0.0
	inverted := false
	perfectMatches := make(map[string]bool)
	for character, matrix := range characterMap {
		var score float64
		switch method {
		case 1:
			score = calculateMSE(testMatrix, matrix)
		case 2:
			score = calculateABS(testMatrix, matrix)
		default:
			score = calculateContrast(testMatrix, matrix, method)
		}

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
	if len(perfectMatches) > 1 {
		// choose randomly
	}
	return match, inverted
}

// getFontBytes Get font btyes from file
func getFontBytes(fontName string) []byte {
	fontBytes, err := ioutil.ReadFile(fontName)
	if err != nil {
		log.Fatalf("ReadFile: %v, %s", err, fontName)
	}

	return fontBytes
}

// RenderText to image
func renderText(text string, fontBytes []byte, fontSize, imageDPI float64, foreground, background *image.Uniform) (image.Image, fixed.Point26_6) {

	// Initialize the context.
	render := image.NewRGBA(image.Rect(0, 0, int(fontSize)*len(text)*2, int(fontSize*2.0)))
	draw.Draw(render, render.Bounds(), background, image.ZP, draw.Src)

	draw.Draw(render, render.Bounds(), background, image.ZP, draw.Src)

	fontData, err := opentype.Parse(fontBytes)
	if err != nil {
		log.Fatalf("Parse: %v, %s", err, *fontName)
	}

	face, err := opentype.NewFace(fontData, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     imageDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("NewFace: %v, %s", err, fontName)
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

	return render, fixed.P(int(bounds.Min.X/64), int(bounds.Max.Y/64))

}

// cropImageToDimension Crops image to the dimensions (x1, y1) to (x2, y2)
func cropImageToDimension(img image.Image, x1, y1, x2, y2 int) image.Image {
	// Calculate the width and height of the cropped region
	width := x2 - x1
	height := y2 - y1

	// Create a new RGBA image with the specified dimensions
	cropped := image.NewRGBA(image.Rect(0, 0, width, height))

	// Copy the cropped region from the original image to the new image
	draw.Draw(cropped, cropped.Bounds(), img, image.Point{x1, y1}, draw.Src)

	return cropped
}

// cropBlank Function to crop the image and remove blank space
func cropBlank(img image.Image) image.Image {
	bounds := img.Bounds()
	minX, minY, maxX, maxY := bounds.Max.X, bounds.Max.Y, 0, 0

	// Find the minimum and maximum non-empty pixel coordinates
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			if _, _, _, a := img.At(x, y).RGBA(); a != 0 {
				if x < minX {
					minX = x
				}
				if x > maxX {
					maxX = x
				}
				if y < minY {
					minY = y
				}
				if y > maxY {
					maxY = y
				}
			}
		}
	}

	// Create a new cropped image
	croppedImage := image.NewRGBA(image.Rect(0, 0, maxX-minX, maxY-minY+10))
	draw.Draw(croppedImage, croppedImage.Bounds(), img, image.Point{minX, minY}, draw.Src)
	return croppedImage
}

// scaleImageToDimension Scales image to the dimensions
func scaleImageToDimension(img image.Image, x, y, method int) image.Image {
	scaled := image.NewRGBA(image.Rect(0, 0, int(x), int(y)))

	switch method {
	case 1:
		draw.ApproxBiLinear.Scale(scaled, scaled.Rect, img, img.Bounds(), draw.Over, nil)
	case 2:
		draw.NearestNeighbor.Scale(scaled, scaled.Rect, img, img.Bounds(), draw.Over, nil)
	case 3:
		draw.CatmullRom.Scale(scaled, scaled.Rect, img, img.Bounds(), draw.Over, nil)
	default:
		draw.BiLinear.Scale(scaled, scaled.Rect, img, img.Bounds(), draw.Over, nil)
	}
	return scaled
}

// scaleImageToProportions scales the image to the size while maintaining proportions
func scaleImageToProportions(img image.Image, size, method int) image.Image {
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
func mapCharacters(characters string, resolution int) map[string][][]int {
	characterMap := make(map[string][][]int)

	for _, character := range characters {
		// fontData := getFontData(fontName)
		// _ = fontData
		renderCharacter, _ := renderText(string(character), gomono.TTF, 20, 72, image.Black, image.Transparent)
		croppedCharacter := cropImageToDimension(renderCharacter, 0, 5, 14, 26)
		croppedCharacter = scaleImageToProportions(croppedCharacter, resolution, 3)
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

// calculateMSE Compares 2 matrix based on a Mean Squared Error
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

// calculateABS Compares 2 matrix based on a average of differences in absolute value
func calculateABS(matrix1, matrix2 [][]int) float64 {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		panic("Matrices must have the same dimensions")
	}

	sum := 0.0
	n := float64(len(matrix1) * len(matrix1[0]))

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[0]); j++ {
			diff := float64(matrix1[i][j] - matrix2[i][j])
			sum += math.Abs(diff)
		}
	}

	return sum / n
}

// calculateContrast Compares 2 matrix based on contrast values
func calculateContrast(matrix1, matrix2 [][]int, thresholdValue int) float64 {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		panic("Matrices must have the same dimensions")
	}

	sum := 0.0
	n := float64(len(matrix1) * len(matrix1[0]))

	for i := 0; i < len(matrix1); i++ {
		for j := 0; j < len(matrix1[0]); j++ {
			score := math.Abs(contrastThreshold(matrix1[i][j], thresholdValue) - contrastThreshold(matrix2[i][j], thresholdValue))
			sum += score
		}
	}
	return sum / n * 255
}

// contrastThreshold Compares a number (contrast) and return 0 if its over it
func contrastThreshold(value, thresholdValue int) float64 {
	if value > thresholdValue {
		return 1.0
	}
	return 0.0
}

// isValidFilePath Validates that a complete path exists and is a file
func isValidFilePath(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		// If os.Stat returns an error, the file or directory doesn't exist
		return false
	}

	return !info.IsDir()
}

// isValidSavePath Validates that the parent directory of a path exists
func isValidSavePath(path string) bool {
	parentDir := filepath.Dir(path)
	info, err := os.Stat(parentDir)
	if err != nil {
		// Error occurred while checking the parent directory
		return false
	}

	// Check if the parent directory exists and is a directory
	return info.IsDir()
}

// outputToFile Writes a character to a file
func outputToFile(filename, character string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(character)
	if err != nil {
		return err
	}

	return writer.Flush()
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

// removeSpecialCharactersAndDuplicates Removes characters the repeat or below 32 in the ascii table
func removeSpecialCharactersAndDuplicates(input string) *string {
	// Create a set to keep track of seen characters
	seen := make(map[rune]bool)

	// Iterate through the string and filter out special characters and duplicates
	var result strings.Builder
	for _, ch := range input {
		if ch >= 32 && !seen[ch] {
			result.WriteRune(ch)
			seen[ch] = true
		}
	}
	filtered := result.String()

	return &filtered
}
