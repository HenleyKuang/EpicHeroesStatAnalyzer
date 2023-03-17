package imageutils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"os"
	"regexp"
)

var (
	heroNameY1Ratio  float32 = 40 / 862.0
	heroNameX2Ratio  float32 = 150 / 388.0
	heroNameY2Ratio  float32 = 60 / 862.0
	mainStatsX1Ratio float32 = 250 / 388.0
	mainStatsY1Ratio float32 = 180 / 862.0
	mainStatsX2Ratio float32 = 330 / 388.0
	mainStatsY2Ratio float32 = 300 / 862.0
	pctStatsX1Ratio  float32 = 250 / 388.0
	pctStatsY1Ratio  float32 = 290 / 862.0
	pctStatsX2Ratio  float32 = 308 / 388.0
	pctStatsY2Ratio  float32 = 600 / 862.0
)

// SubImager is an interface for cropping the image.
type SubImager interface {
	SubImage(r image.Rectangle) image.Image
}

// Base64ToFileObject converts a base64 string representation of an iage to a file object.
func Base64ToFileObject(imgAsBase64 string) (*os.File, error) {
	tempfile, err := ioutil.TempFile("", "ocrserver"+"-")
	if err != nil {
		return nil, err
	}
	defer func() {
		tempfile.Close()
		os.Remove(tempfile.Name())
	}()

	if len(imgAsBase64) == 0 {
		return nil, fmt.Errorf("base64 string required")
	}
	imgAsBase64 = regexp.MustCompile("data:image\\/png;base64,").ReplaceAllString(imgAsBase64, "")
	b, err := base64.StdEncoding.DecodeString(imgAsBase64)
	if err != nil {
		return nil, err
	}
	tempfile.Write(b)
	return tempfile, nil
}

// FileObjectToBytes converts an image file object to bytes.
func FileObjectToBytes(imgFile *os.File) ([]byte, error) {
	buff := make([]byte, 512)
	_, err := imgFile.Read(buff)
	if err != nil {
		return nil, err
	}
	return buff, nil
}

// BytesToImageObject converts an image represented as bytes to an image object.
func BytesToImageObject(imgBytes []byte) (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(imgBytes))
	if err != nil {
		return nil, err
	}
	return img, nil
}

// ImageObjectToBytes converts an image object to bytes.
func ImageObjectToBytes(imgObj image.Image) ([]byte, error) {
	// Create buffer.
	buff := new(bytes.Buffer)

	// Encode image to buffer.
	err := png.Encode(buff, imgObj)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

// FileObjectToImageObject converts a file object to an image object.
func FileObjectToImageObject(imgFile *os.File) (image.Image, error) {
	img, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return img, nil
}

// ImageObjToGrayScale grayscales a given image obj.
func ImageObjToGrayScale(imgObj image.Image) *image.Gray {
	var (
		bounds = imgObj.Bounds()
		gray   = image.NewGray(bounds)
	)
	for x := 0; x < bounds.Max.X; x++ {
		for y := 0; y < bounds.Max.Y; y++ {
			var rgba = imgObj.At(x, y)
			gray.Set(x, y, rgba)
		}
	}
	return gray
}

// CropImage crops a given image to a specific Rect.
func CropImage(imgObj image.Image, crop image.Rectangle) (image.Image, error) {
	simg, ok := imgObj.(SubImager)
	if !ok {
		return nil, fmt.Errorf("image does not support cropping")
	}
	return simg.SubImage(crop), nil
}

// CropToHeroName crops a given image to the section that contains the Hero Name.
func CropToHeroName(imgObj image.Image) (image.Image, error) {
	w := float32(imgObj.Bounds().Dx())
	h := float32(imgObj.Bounds().Dy())
	// cropHeroName := image.Rect(0, 40, 150, 60)
	cropHeroName := image.Rect(0, int(heroNameY1Ratio*h), int(heroNameX2Ratio*w), int(heroNameY2Ratio*h))
	return CropImage(imgObj, cropHeroName)
}

// CropToMainStats crops a given image to the section that contains the main stats (non-percentage stats)
func CropToMainStats(imgObj image.Image) (image.Image, error) {
	w := float32(imgObj.Bounds().Dx())
	h := float32(imgObj.Bounds().Dy())
	// cropMainStats := image.Rect(250, 180, 330, 300)
	cropMainStats := image.Rect(int(mainStatsX1Ratio*w), int(mainStatsY1Ratio*h), int(mainStatsX2Ratio*w), int(mainStatsY2Ratio*h))
	return CropImage(imgObj, cropMainStats)
}

// CropToPercentageStats crops a given image to the section that contains the stats with percentages.
func CropToPercentageStats(imgObj image.Image) (image.Image, error) {
	w := float32(imgObj.Bounds().Dx())
	h := float32(imgObj.Bounds().Dy())
	// cropPercentageStats := image.Rect(250, 290, 308, 600)
	cropPercentageStats := image.Rect(int(pctStatsX1Ratio*w), int(pctStatsY1Ratio*h), int(pctStatsX2Ratio*w), int(pctStatsY2Ratio*h))
	return CropImage(imgObj, cropPercentageStats)
}
