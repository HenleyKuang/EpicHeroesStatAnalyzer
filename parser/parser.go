package parser

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/otiai10/gosseract/v2"
)

var mainStatsOrder = []string{"Power", "HP", "ATK", "Armor", "Speed"}
var percentageStatsOrder = []string{"Crit", "Crit Resistance", "Crit DMG", "Crit Damage Resistance", "Skill DMG", "Holy DMG",
	"Effect Hit", "Effect Res", "Hit", "Dodge", "Accuracy", "Block", "Broken Armor", "DMG Immune"}

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
func MainStatsFromBytes(client *gosseract.Client, imgAsBytes []byte) (map[string]int, error) {
	client.SetImageFromBytes(imgAsBytes)
	stats, err := client.Text()
	if err != nil {
		return nil, err
	}
	statsSplit := strings.Split(stats, "\n")
	if len(statsSplit) != 5 {
		return nil, fmt.Errorf("Expected 5 stats to be parsed. Got %d", len(statsSplit))
	}
	statsMap := map[string]int{}
	for idx, statName := range mainStatsOrder {
		statValue, _ := strconv.Atoi(statsSplit[idx])
		statsMap[statName] = statValue
	}
	return statsMap, nil
}

// PercentageStatsFromBytes returns the list of the percentage stats given the image in byte format.
func PercentageStatsFromBytes(client *gosseract.Client, imgAsBytes []byte) (map[string]float32, error) {
	client.SetImageFromBytes(imgAsBytes)
	stats, err := client.Text()
	if err != nil {
		return nil, err
	}
	statsSplit := strings.Split(stats, "\n")
	if len(statsSplit) != 14 {
		return nil, fmt.Errorf("Expected 14 stats to be parsed. Got %d", len(statsSplit))
	}
	statsMap := map[string]float32{}
	for idx, statName := range percentageStatsOrder {
		statValue, _ := strconv.ParseFloat(statsSplit[idx], 32)
		statsMap[statName] = float32(statValue)
	}
	return statsMap, nil
}
