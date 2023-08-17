package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"flag"
	"fmt"
	"golang.org/x/image/draw"
	"golang.org/x/image/font"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/png"
	"io/ioutil"
	"log"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var (
	// Render options
	text              = flag.String("text", "Hello World!", "text to render")
	displayCharacters = flag.String("characters", "", "text to render, ignored when loading a map")
	displayResolution = flag.Int("resolution", 16, "text to render, ignored when loading a map")
	pixelAspect       = flag.Float64("aspect", 0.5, "character height to width")
	fontName          = flag.String("font", getDefaultFont(), "filename of the ttf/otf font")
	fontSize          = flag.Float64("size", 25.0, "font size in points")
	maxWidth          = flag.Int("max-width", 0, "maximum width to render")
	useInverted       = flag.Bool("allow-inverted", false, "use inverted characters, ignored when writing to file")
	inverted          = flag.Bool("invert", false, "invert image")
	renderMode        = flag.Int("mode", 20, "render mode")

	// File options
	loadFile   = flag.String("load", "", "load saved character map")
	saveFile   = flag.String("save", "", "save character map")
	outputFile = flag.String("output", "", "save output to file")

	debug = flag.Bool("debug", false, "debug mode")
)

func main() {
	// Custom help usage
	flag.Usage = func() {
		flagSet := flag.CommandLine
		fmt.Printf("Usage: %s\n\n", filepath.Base(flagSet.Name()))

		//goland:noinspection GoPrintFunctions
		fmt.Println("Create a title\n")

		fmt.Println("optional arguments:")
		order := map[string][]string{
			"display":      {"text", "characters", "resolution", "aspect", "font", "size", "max-width", "inverted", "mode", "allow-inverted"},
			"input/output": {"load", "save", "output"},
		}
		for key, values := range order {
			fmt.Println(key + ":")
			for _, value := range values {
				flagOption := flagSet.Lookup(value)
				fmt.Printf("    --%s", flagOption.Name)
				if flagOption.Value.String() == "0" || flagOption.Value.String() == "" {
					fmt.Printf("=%s", "None")
				} else {
					fmt.Printf("=%s", flagOption.Value)
				}
				fmt.Println()
				fmt.Printf("        %s ", flagOption.Usage)
				if flagOption.DefValue != "0" && flagOption.DefValue != "" && flagOption.DefValue != "false" {
					fmt.Printf("(default %s)", flagOption.DefValue)
				}
				fmt.Println()
			}
		}
	}
	flag.Parse()

	// Process any arguments that new validation or alteration
	if *displayCharacters == "" {
		for x := 32; x < 128; x++ {
			*displayCharacters += string(rune(x))
		}
		for x := 161; x < 173; x++ {
			*displayCharacters += string(rune(x))
		}
		for x := 174; x < 255; x++ {
			*displayCharacters += string(rune(x))
		}
	} else {
		displayCharacters = removeSpecialCharactersAndDuplicates(*displayCharacters)
	}

	if !isValidFilePath(*fontName) {
		fmt.Printf("%s does not exist", *fontName)
		os.Exit(1)
	}

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
	var err error
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
			characterMap, err = loadCharacterMapFromDisk(*loadFile)
			if err != nil {
				log.Fatal(err)
			}
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
	fontBytes := getFont(*fontName)
	renderImage, textImageRect := renderText("\n"+*text, *fontBytes, *fontSize, 72.0, image.White, image.Transparent)
	renderImage = cropImageToDimension(renderImage, 0, 0, int(textImageRect.X>>6), int(textImageRect.Y>>6))

	// Check if the image width is greater than window width and if we are rending to screen
	if *maxWidth == 0 {
		if screenOutput {
			winSize, err := getWinSize()
			if err != nil {
				fmt.Println(err)
				*maxWidth = math.MaxInt32
			}
			*maxWidth = winSize
		} else {
			*maxWidth = math.MaxInt32
		}
	}
	displayWidth := *maxWidth * *displayResolution
	if int(textImageRect.X>>6) > displayWidth {
		// Calculate the scale factor to resize the image to width 120
		scale := float64(displayWidth) / float64(textImageRect.X>>6)
		renderImage = scaleImageToDimension(renderImage, displayWidth, int(float64(textImageRect.Y>>6)*scale), 0)
	}
	if *debug {
		err := saveImage(renderImage, "title.png")
		if err != nil {
			log.Println(err)
		}
	}

	// Scale to pixel ratio
	bounds := renderImage.Bounds()
	renderImage = scaleImageToDimension(renderImage, bounds.Max.X, int(float64(bounds.Max.Y)**pixelAspect), 0)
	renderImage = cropImageToContent(renderImage, *displayResolution)

	// Build and fill a 4D matrix of the title
	brightnessMatrix := getEmptyBrightnessMatrix(renderImage, *displayResolution)
	brightnessMatrix = fillBrightnessMatrix(brightnessMatrix, renderImage, *displayResolution)

	// Build character map if not loading from file
	if !isValidFilePath(*loadFile) {
		characterMap = mapCharacters(*displayCharacters, *displayResolution)
	}

	// Save character map
	if *saveFile != "" && isValidSavePath(*saveFile) {
		answer := ""
		if isValidFilePath(*saveFile) {
			for {
				fmt.Printf("File %s already exists. Do you want to replace it? (yes/no): ", *saveFile)

				reader := bufio.NewReader(os.Stdin)
				answer, _ = reader.ReadString('\n')
				answer = strings.TrimSpace(answer)

				if answer == "no" || answer == "n" || answer == "yes" || answer == "y" {
					break
				}
			}
		}
		if answer == "yes" || answer == "y" || answer == "" {
			if strings.HasSuffix(*saveFile, ".json") {
				fmt.Println("Save json to disk")

				// Open a file for writing
				file, err := os.Create(*saveFile)
				if err != nil {
					log.Fatal(err)
				}
				defer func(file *os.File) {
					err := file.Close()
					if err != nil {
					}
				}(file)

				// Create a JSON encoder to write to the file
				encoder := json.NewEncoder(file)

				// Encode the characterMap and write it to the file
				err = encoder.Encode(characterMap)
				if err != nil {
					log.Fatal(err)
				}
			} else {
				fmt.Println("Save map to disk")
				if err := saveCharacterMapToDisk(characterMap, *saveFile); err != nil {
					log.Fatal(err)
				}
			}
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
			character, invert := findMatch(characterMap, brightnessMatrix[x][y], *renderMode, *inverted)
			if invert {
				fmt.Print("\x1B[7m")
			}
			if screenOutput {
				fmt.Print(character)
			} else {
				if err := outputToFile(*outputFile, character); err != nil {
					log.Fatal(err)
				}
			}
			fmt.Print("\x1B[0m")
		}
		if screenOutput {
			fmt.Println()
		} else {
			if err := outputToFile(*outputFile, "\n"); err != nil {
				log.Fatal(err)
			}
		}
	}
	if !screenOutput {
		fmt.Printf("Output added to file: %s", *outputFile)
	}
}

// findMatch Find a character that matches a section of image
func findMatch(characterMap map[string][][]int, testMatrix [][]int, method int, inverted bool) (string, bool) {
	match := " "
	bestScore := 255.0
	invertedMatch := " "
	invertedBestScore := 0.0
	invertCharacters := false
	for character, matrix := range characterMap {
		var score float64
		switch method {
		case 1:
			score = calculateABS(testMatrix, matrix)
		case 2:
			score = compareNeighbours(testMatrix, matrix, 1)
		case 3:
			score = compareNeighbours(testMatrix, matrix, 2)
		case 4:
			score = compareNeighbours(testMatrix, matrix, 3)
		case 10:
			score = calculateContrast(testMatrix, matrix, 1)
		case 11:
			score = calculateContrast(testMatrix, matrix, 2)
		case 12:
			score = calculateContrast(testMatrix, matrix, 4)
		case 13:
			score = calculateContrast(testMatrix, matrix, 8)
		case 14:
			score = calculateContrast(testMatrix, matrix, 16)
		case 15:
			score = calculateContrast(testMatrix, matrix, 32)
		case 16:
			score = calculateContrast(testMatrix, matrix, 64)
		case 17:
			score = calculateContrast(testMatrix, matrix, 128)
		case 18:
			score = calculateContrast(testMatrix, matrix, 192)
		case 19:
			score = calculateContrast(testMatrix, matrix, 224)
		case 20:
			score = calculateContrast(testMatrix, matrix, 240)
		case 21:
			score = calculateContrast(testMatrix, matrix, 248)
		case 22:
			score = calculateContrast(testMatrix, matrix, 252)
		case 23:
			score = calculateContrast(testMatrix, matrix, 254)
		default:
			score = calculateMSE(testMatrix, matrix)
		}

		if inverted {
			score = 255 - score
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
			invertCharacters = false
		} else {
			match = invertedMatch
			invertCharacters = true
		}
	}
	return match, invertCharacters
}

// getFont Get font from file
func getFont(fontName string) *sfnt.Font {
	fontBytes, err := ioutil.ReadFile(fontName)
	if err != nil {
		log.Fatalf("ReadFile: %v, %s", err, fontName)
	}

	switch filepath.Ext(fontName) {
	case ".dfont":
		fallthrough
	case ".ttc":
		fallthrough
	case ".otc":
		collection, err := sfnt.ParseCollection(fontBytes)
		if err != nil {
			log.Fatal(err)
		}
		fontData, err := collection.Font(0)
		if err != nil {
			log.Fatalf("Font collection: %v", err)
		}

		return fontData
	case ".ttf":
		fallthrough
	case ".otf":
		fallthrough
	case "":
		fontData, err := opentype.Parse(fontBytes)
		if err != nil {
			log.Fatalf("Parse: %v, %s", err, fontName)
		}
		return fontData
	}
	return nil
}

// renderText to image
func renderText(text string, fontData sfnt.Font, fontSize, imageDPI float64, foreground, background *image.Uniform) (image.Image, fixed.Point26_6) {

	// Initialize the context.
	lines := strings.Split(text, "\n")
	render := image.NewRGBA(image.Rect(0, 0, int(fontSize)*len(text)*2, int(fontSize)*(len(lines)+1)))
	draw.Draw(render, render.Bounds(), background, image.Pt(0, 0), draw.Src)

	face, err := opentype.NewFace(&fontData, &opentype.FaceOptions{
		Size:    fontSize,
		DPI:     imageDPI,
		Hinting: font.HintingFull,
	})
	if err != nil {
		log.Fatalf("NewFace: %v, %s", err, *fontName)
	}
	fontDrawer := font.Drawer{
		Dst:  render,
		Src:  foreground,
		Face: face,
		Dot:  fixed.P(0, int(fontSize)),
	}
	maxDot := fixed.I(0)
	for _, line := range lines {
		fontDrawer.Dot.X = 0
		fontDrawer.DrawString(line)
		if fontDrawer.Dot.X > maxDot {
			maxDot = fontDrawer.Dot.X
		}
		fontDrawer.Dot.Y += fixed.I(int(fontSize))

	}
	fontDrawer.Dot.X = maxDot
	// Determine the rendering bounds of the text
	bounds, _ := fontDrawer.BoundString(text)

	return render, fixed.P(int(bounds.Min.X/64), int(bounds.Max.Y/64))

}

// cropImageToContent Find the first and last line with content and crop
func cropImageToContent(imageSrc image.Image, resolution int) image.Image {
	bounds := imageSrc.Bounds()
	startY, endY := -1, -1

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			_, _, _, a := imageSrc.At(x, y).RGBA()
			if a > 0 {
				if startY < 0 {
					startY = y
					break
				}
				if startY >= 0 {
					endY = y
					break
				}
			}
		}
	}

	return cropImageToDimension(imageSrc, bounds.Min.X, startY, bounds.Max.X, endY+resolution)
}

// cropImageToDimension Crops image to the dimensions (x1, y1) to (x2, y2)
func cropImageToDimension(img image.Image, x1, y1, x2, y2 int) image.Image {
	// Calculate the width and height of the cropped region
	width := x2 - x1
	height := y2 - y1

	// Create a new RGBA image with the specified dimensions
	cropped := image.NewRGBA(image.Rect(0, 0, width, height))

	// Copy the cropped region from the original image to the new image
	draw.Draw(cropped, cropped.Bounds(), img, image.Point{X: x1, Y: y1}, draw.Src)

	return cropped
}

// scaleImageToDimension Scales image to the dimensions
func scaleImageToDimension(img image.Image, x, y, method int) image.Image {
	scaled := image.NewRGBA(image.Rect(0, 0, x, y))

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
	draw.Draw(resized, resized.Bounds(), img, image.Point{X: offsetX, Y: offsetY}, draw.Src)

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

	// Create the 2D matrix of res X res matrices
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

// fillBrightnessMatrix reads an image file, and it fills a matrix of with brightness values.
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

// mapCharacters Create a brightness map of the characters
func mapCharacters(characters string, resolution int) map[string][][]int {
	characterMap := make(map[string][][]int)

	for _, character := range characters {
		fontData, err := sfnt.Parse(gomono.TTF)
		if err != nil {
			log.Fatalf("Parse: %v, %s", err, *fontName)
		}
		renderCharacter, _ := renderText(string(character), *fontData, 20, 72, image.Black, image.Transparent)
		croppedCharacter := cropImageToDimension(renderCharacter, 0, 5, 14, 26)
		croppedCharacter = scaleImageToProportions(croppedCharacter, resolution, 3)
		if *debug {
			err := saveImage(croppedCharacter, fmt.Sprintf("%d.png", character))
			if err != nil {
				log.Println(err)
			}
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
			brightness := uint32(rgba.A)

			// Store the brightness value in the matrix
			characterMatrix[x][y] = int(brightness)
		}
	}

	return characterMatrix
}

// saveImage Debugging, save image to disk
func saveImage(img image.Image, name string) error {
	output, _ := os.Create(name)
	defer func(output *os.File) {
		err := output.Close()
		if err != nil {
		}
	}(output)

	err := png.Encode(output, img)
	if err != nil {
		return err
	}
	fmt.Printf("Image saved to %s\n", name)

	return nil
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
		panic("Matrices must have the same dimensions.")
	}

	rows, cols := len(matrix1), len(matrix1[0])
	sum := 0.0

	for x := 0; x < rows; x++ {
		for y := 0; y < cols; y++ {
			diff := matrix1[x][y] - matrix2[x][y]
			sum += float64(diff * diff)
		}
	}

	mse := sum / float64(rows*cols)
	return math.Sqrt(mse) // The score ranges from 0 (match) to 255 (exact opposite)
}

// calculateABS Compares 2 matrix based on an average of differences in absolute value
func calculateABS(matrix1, matrix2 [][]int) float64 {
	if len(matrix1) != len(matrix2) || len(matrix1[0]) != len(matrix2[0]) {
		panic("Matrices must have the same dimensions")
	}

	sum := 0.0
	n := float64(len(matrix1) * len(matrix1[0]))

	for x := 0; x < len(matrix1); x++ {
		for y := 0; y < len(matrix1[0]); y++ {
			diff := float64(matrix1[x][y] - matrix2[x][y])
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

	for x := 0; x < len(matrix1); x++ {
		for y := 0; y < len(matrix1[0]); y++ {
			score := math.Abs(contrastThreshold(matrix1[x][y], thresholdValue) - contrastThreshold(matrix2[x][y], thresholdValue))
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

// compareNeighbours Compares 2 matrix with each point the nearest neighbours are wighted too
func compareNeighbours(matrix1, matrix2 [][]int, n int) float64 {

	totalDifference := 0.0
	numElements := 0

	// Loop through each element in the matrices
	for x := 0; x < len(matrix1); x++ {
		for y := 0; y < len(matrix1[0]); y++ {
			// Calculate the sum of absolute differences with neighbors
			differenceSum := 0.0
			count := 0

			for nx := -n; nx <= n; nx++ {
				for ny := -n; ny <= n; ny++ {
					if x+nx >= 0 && x+nx < len(matrix1) && y+ny >= 0 && y+ny < len(matrix1[0]) {
						differenceSum += math.Abs(float64(matrix1[x+nx][y+ny] - matrix2[x+nx][y+ny]))
						count++
					}
				}
			}

			// Update the total difference
			if count > 0 {
				averageDifference := differenceSum / float64(count)
				totalDifference += averageDifference
				numElements++
			}
		}
	}

	// Calculate the similarity score as the average difference
	// normalized to a range of 0-255
	averageDifference := totalDifference / float64(numElements)
	return averageDifference
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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

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
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
		}
	}(file)

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

// getDefaultFont Get the default font for each OS
func getDefaultFont() string {
	switch runtime.GOOS {
	case "darwin": // macOS
		return "/System/Library/Fonts/Helvetica.ttc"
	case "windows":
		return "C:\\Windows\\Fonts\\arial.ttf"
	case "linux":
		return "/usr/share/fonts/truetype/dejavu/DejaVuSans.ttf"
	default:
		return ""
	}
}
