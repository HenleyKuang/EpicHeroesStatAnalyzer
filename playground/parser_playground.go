package main

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"log"
	"main/imageutils"
	"main/parser"
	"os"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	// Read the entire file into a byte slice
	imgPath := "./playground/data/toko_stats.jpg"
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
	imgObj, err := imageutils.BytesToImageObject(bytes)
	if err != nil {
		fmt.Println("[BytesToImageObject] ", err)
		os.Exit(3)
	}
	croppedImgObj, err := imageutils.CropImage(imgObj, image.Rect(0, 40, 150, 60))
	if err != nil {
		fmt.Println("[CropImage] ", err)
		os.Exit(3)
	}
	croppedImgBytes, err := imageutils.ImageObjectToBytes(croppedImgObj)
	if err != nil {
		fmt.Println("[ImageObjectToBytes] ", err)
		os.Exit(3)
	}
	// Save cropped image.
	out, _ := os.Create("./playground/data/toko_stats_cropped.jpg")
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 100
	err = jpeg.Encode(out, croppedImgObj, &opts)
	//jpeg.Encode(out, img, nil)
	if err != nil {
		log.Println("[jpeg.Encode] ", err)
	}
	// Get Hero Name.
	heroName, err := parser.HeroNameFromBytes(croppedImgBytes)
	// heroName, err := parser.HeroNameFromBytes(bytes)

	// var base64Encoding string

	// // Determine the content type of the image file
	// mimeType := http.DetectContentType(bytes)

	// // Prepend the appropriate URI scheme header depending
	// // on the MIME type
	// switch mimeType {
	// case "image/jpeg":
	// 	base64Encoding += "data:image/jpeg;base64,"
	// case "image/png":
	// 	base64Encoding += "data:image/png;base64,"
	// }

	// // Append the base64 encoded output
	// base64Encoding += toBase64(bytes)

	// // Print the full base64 representation of the image
	// fmt.Println(base64Encoding)

	// heroName, err := parser.HeroNameFromBase64(base64Encoding); err != nil {
	fmt.Println("[HeroNameFromBytes] ", err)
	fmt.Println(heroName)

	// client := gosseract.NewClient()
	// defer client.Close()

	// client.SetWhitelist("0123456789")
	// client.SetImage("./playground/data/toko_stats.jpg")
	// heroName, err := client.Text()
	// fmt.Println(err)
	// fmt.Println(heroName)
}
