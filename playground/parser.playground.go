package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"main/parser"
)

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {
	// Read the entire file into a byte slice
	bytes, err := ioutil.ReadFile("./playground/data/toko_stats.jpg")
	if err != nil {
		log.Fatal(err)
	}

	heroName, err := parser.HeroNameFromBytes(bytes)

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
	fmt.Println(err)
	fmt.Println(heroName)

	// client := gosseract.NewClient()
	// defer client.Close()

	// client.SetWhitelist("0123456789")
	// client.SetImage("./playground/data/toko_stats.jpg")
	// heroName, err := client.Text()
	// fmt.Println(err)
	// fmt.Println(heroName)
}
