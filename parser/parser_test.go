package parser

import (
	"bufio"
	"fmt"
	"main/imageutils"
	"os"
	"strings"
	"testing"
)

func readFileToBytes(imgPath string) []byte {
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
		return nil
	}
	return bytes
}

func TestHeroNameFromBytes(t *testing.T) {
	want := "Samurai Girl"
	bytes := readFileToBytes("./data/toko_stats.jpg")
	imgObj, _ := imageutils.BytesToImageObject(bytes)
	croppedImgObj, _ := imageutils.CropToHeroName(imgObj)
	croppedImgBytes, _ := imageutils.ImageObjectToBytes(croppedImgObj)
	got, err := HeroNameFromBytes(croppedImgBytes)
	if !strings.HasPrefix(got, want) || err != nil {
		t.Fatalf(`HeroNameFromBytes() = %q, %v, want match for %#q, nil`, got, err, want)
	}
}
