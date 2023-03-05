package parser

import (
	"bufio"
	"main/image"

	"github.com/otiai10/gosseract/v2"
)

func init() {
	client := gosseract.NewClient()
	defer client.Close()
}

// HeroNameFromBase64 returns the hero's name given the image in base64 format.
func HeroNameFromBase64(imgAsBase64 string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	imgFile, err := image.FromBase64(imgAsBase64)
	if err != nil {
		return "", err
	}
	defer imgFile.Close()

	// Convert img file to bytes.
	fileInfo, _ := imgFile.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	// Read file into bytes.
	buffer := bufio.NewReader(imgFile)
	_, err = buffer.Read(bytes)

	if err != nil {
		return "", err
	}

	client.SetWhitelist("0123456789")
	client.SetImageFromBytes(bytes)

	heroName, _ := client.Text()
	return heroName, nil
}

// HeroNameFromBytes returns the hero's name given the image in byte format.
func HeroNameFromBytes(imgAsBytes []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	client.SetWhitelist("0123456789")
	client.SetImageFromBytes(imgAsBytes)

	heroName, _ := client.Text()
	return heroName, nil
}
