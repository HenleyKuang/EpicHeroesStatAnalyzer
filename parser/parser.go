package parser

import (
	"bufio"
	"main/imageutils"

	"github.com/otiai10/gosseract/v2"
)

// HeroNameFromBase64 returns the hero's name given the image in base64 format.
func HeroNameFromBase64(imgAsBase64 string) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	imgFile, err := imageutils.Base64ToFileObject(imgAsBase64)
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
func HeroNameFromBytes(client *gosseract.Client, imgAsBytes []byte) (string, error) {
	client.SetImageFromBytes(imgAsBytes)
	heroName, err := client.Text()
	if err != nil {
		return "", err
	}
	return heroName, nil
}

// MainStatsFromBytes returns the list of the main stats (first 5 non-percentage stats) given the image in byte format.
func MainStatsFromBytes(imgAsBytes []byte) (string, error) {
	client := gosseract.NewClient()
	client.SetTessdataPrefix("./traineddata/")
	client.SetLanguage("eng")
	defer client.Close()

	client.SetWhitelist("0123456789%.")
	client.SetImageFromBytes(imgAsBytes)
	stats, err := client.Text()
	if err != nil {
		return "", err
	}
	return stats, nil
}
