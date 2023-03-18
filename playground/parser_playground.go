package main

import (
	"bufio"
	"fmt"
	"image/jpeg"
	"log"
	"main/imageutils"
	"main/parser"
	"os"

	"github.com/otiai10/gosseract/v2"
)

func readFileToBytes() []byte {
	// Read the entire file into a byte slice
	// imgPath := "./playground/data/toko_stats.jpg"
	imgPath := "./playground/data/indira_stats.jpg"
	fileObj, err := os.Open(imgPath)
	if err != nil {
		fmt.Println("os.Open ", err)
		os.Exit(1)
	}
	defer fileObj.Close()

	fileInfo, _ := fileObj.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)
	// read file into bytes
	buffer := bufio.NewReader(fileObj)
	_, err = buffer.Read(bytes)
	if err != nil {
		fmt.Println("[buffer.Read] ", err)
		os.Exit(1)
	}
	return bytes
}

func main() {
	bytes := readFileToBytes()
	imgObj, err := imageutils.BytesToImageObject(bytes)
	if err != nil {
		fmt.Println("[BytesToImageObject] ", err)
		os.Exit(3)
	}
	imgObj = imageutils.ImageObjToGrayScale(imgObj)
	// croppedImgObj, err := imageutils.CropToHeroName(imgObj)
	croppedImgObj, err := imageutils.CropToMainStats(imgObj)
	// croppedImgObj, err := imageutils.CropToPercentageStats(imgObj)
	if err != nil {
		fmt.Println("[CropToHeroName] ", err)
		os.Exit(3)
	}
	croppedImgBytes, err := imageutils.ImageObjectToBytes(croppedImgObj)
	if err != nil {
		fmt.Println("[ImageObjectToBytes] ", err)
		os.Exit(3)
	}
	// Save cropped image.
	// out, _ := os.Create("./playground/data/toko_main_stats_cropped.jpg")
	out, _ := os.Create("./playground/data/indira_main_stats_cropped.jpg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100
	err = jpeg.Encode(out, croppedImgObj, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println("[jpeg.Encode] ", err)
	}
	// Get Hero Name.
	// alphabetClient := gosseract.NewClient()
	// alphabetClient.SetTessdataPrefix("./traineddata/")
	// alphabetClient.SetLanguage("eng")
	// alphabetClient.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ")
	// defer alphabetClient.Close()
	// heroName, err := parser.HeroNameFromBytes(alphabetClient, croppedImgBytes)

	// Get stats
	digitsClient := gosseract.NewClient()
	digitsClient.SetTessdataPrefix("./traineddata/")
	digitsClient.SetLanguage("digitsall_layer")
	digitsClient.SetWhitelist("0123456789")
	defer digitsClient.Close()
	stats, err := parser.MainStatsFromBytes(digitsClient, croppedImgBytes)
	fmt.Println("[MainStatsFromBytes] ", err)
	fmt.Println(stats)

	// client := gosseract.NewClient()
	// defer client.Close()

	// client.SetWhitelist("0123456789")
	// client.SetImage("./playground/data/toko_stats.jpg")
	// heroName, err := client.Text()
	// fmt.Println(err)
	// fmt.Println(heroName)

	// digitsClient := gosseract.NewClient()
	// digitsClient.SetTessdataPrefix("./traineddata/")
	// digitsClient.SetLanguage("digitsall_layer")
	// digitsClient.SetWhitelist("0123456789.")
	// defer digitsClient.Close()
	// stats, err := parser.PercentageStatsFromBytes(digitsClient, croppedImgBytes)
	// fmt.Println("[PercentageStatsFromBytes] ", err)
	// fmt.Println(stats)
}
