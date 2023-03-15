package parser

import (
	"github.com/otiai10/gosseract/v2"
)

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
func MainStatsFromBytes(client *gosseract.Client, imgAsBytes []byte) (string, error) {
	client.SetImageFromBytes(imgAsBytes)
	stats, err := client.Text()
	if err != nil {
		return "", err
	}
	return stats, nil
}
