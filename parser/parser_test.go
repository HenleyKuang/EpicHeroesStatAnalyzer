package parser

import (
	"bufio"
	"fmt"
	"main/imageutils"
	"os"
	"strings"
	"testing"

	"github.com/otiai10/gosseract/v2"
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
	client := gosseract.NewClient()
	client.SetTessdataPrefix("../traineddata/")
	client.SetLanguage("eng")
	client.SetWhitelist("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz ")
	defer client.Close()

	want := "Samurai Girl"
	bytes := readFileToBytes("./data/toko_stats.jpg")
	imgObj, _ := imageutils.BytesToImageObject(bytes)
	croppedImgObj, _ := imageutils.CropToHeroName(imgObj)
	croppedImgBytes, _ := imageutils.ImageObjectToBytes(croppedImgObj)
	got, err := HeroNameFromBytes(client, croppedImgBytes)
	if !strings.HasPrefix(got, want) || err != nil {
		t.Fatalf(`HeroNameFromBytes() = %q, %v, want match for %#q, nil`, got, err, want)
	}
}

func TestMainStatsFromBytes(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		want     map[string]interface{}
	}{
		{
			name:     "toko main stats",
			fileName: "./data/toko_stats.jpg",
			want: map[string]interface{}{
				"Power": 6278276,
				"HP":    11285601,
				"ATK":   637280,
				"Armor": 4468,
				"Speed": 133,
			},
		},
		{
			name:     "kinley main stats",
			fileName: "./data/kinley_stats.jpeg",
			want: map[string]interface{}{
				"Power": 5912727,
				"HP":    13621736,
				"ATK":   459930,
				"Armor": 4834,
				"Speed": 127,
			},
		},
		{
			name:     "indira main stats",
			fileName: "./data/indira_stats.jpg",
			want: map[string]interface{}{
				"Power": 6360034,
				"HP":    10862119,
				"ATK":   675963,
				"Armor": 4610,
				"Speed": 127,
			},
		},
	}

	client := gosseract.NewClient()
	client.SetTessdataPrefix("../traineddata/")
	client.SetLanguage("digitsall_layer")
	client.SetWhitelist("0123456789")
	defer client.Close()

	for _, test := range tests {
		bytes := readFileToBytes(test.fileName)
		imgObj, _ := imageutils.BytesToImageObject(bytes)
		croppedImgObj, _ := imageutils.CropToMainStats(imgObj)
		croppedImgBytes, _ := imageutils.ImageObjectToBytes(croppedImgObj)
		got, err := MainStatsFromBytes(client, croppedImgBytes)
		if fmt.Sprint(got) != fmt.Sprint(test.want) || err != nil {
			t.Fatalf(`%s: MainStatsFromBytes() = %v, %v, want match for %v, nil`, test.name, got, err, test.want)
		}
	}
}

func TestPercentageStatsFromBytes(t *testing.T) {
	client := gosseract.NewClient()
	client.SetTessdataPrefix("../traineddata/")
	client.SetLanguage("digitsall_layer")
	client.SetWhitelist("0123456789.")
	defer client.Close()

	want := map[string]interface{}{
		"Crit":                   81,
		"Crit Resistance":        4,
		"Crit DMG":               31,
		"Crit Damage Resistance": 15,
		"Skill DMG":              33,
		"Holy DMG":               0,
		"Effect Hit":             24,
		"Effect Res":             28,
		"Hit":                    0,
		"Dodge":                  0,
		"Accuracy":               46,
		"Block":                  5.6,
		"Broken Armor":           65,
		"DMG Immune":             23,
	}

	bytes := readFileToBytes("./data/toko_stats.jpg")
	imgObj, _ := imageutils.BytesToImageObject(bytes)
	croppedImgObj, _ := imageutils.CropToPercentageStats(imgObj)
	croppedImgBytes, _ := imageutils.ImageObjectToBytes(croppedImgObj)
	got, err := PercentageStatsFromBytes(client, croppedImgBytes)
	if fmt.Sprint(got) != fmt.Sprint(want) || err != nil {
		t.Fatalf(`PercentageStatsFromBytes() = %v, %v, want match for %v, nil`, got, err, want)
	}
}
