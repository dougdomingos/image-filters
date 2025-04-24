package utils

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
)

// LoadImage receives the path of an image file and returns it as an RGBA image,
// along with its original format (e.g., "jpeg", "png").
func LoadImage(filePath string) (*image.RGBA, string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, "", fmt.Errorf("[ERROR] Unable to load %s: %w", filePath, err)
	}
	defer file.Close()

	img, format, err := image.Decode(file)
	if err != nil {
		return nil, "", fmt.Errorf("[ERROR] Unable to decode %s: %w", filePath, err)
	}

	bounds := img.Bounds()
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, img, bounds.Min, draw.Src)

	return rgba, format, nil
}

// SaveImage takes a RGBA image, its original encoding format and the path in
// which it'll be stored and creates a new file containing the specified image
// at that path.
func SaveImage(img *image.RGBA, format string, outputDir string, outputFile string) error {
	file, err := os.Create(filepath.Join(outputDir, outputFile))
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to create file \"%s\": %w", outputDir, err)
	}
	defer file.Close()

	var encodingErr error
	switch format {
	case "jpeg":
		encodingErr = jpeg.Encode(file, img, &jpeg.Options{Quality: 95})
	case "png":
		encodingErr = png.Encode(file, img)
	default:
		encodingErr = fmt.Errorf("[ERROR] Unsuported format: %s", format)
	}

	return encodingErr
}

// CreateOutputDir checks if a directory exists at the given path, and creates
// it (with parents) if it doesn't.
func CreateOutputDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return fmt.Errorf("[ERROR]: Creating output directory failed! %s", err)
	}

	return err
}
