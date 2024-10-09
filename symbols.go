package main

import (
	"math/rand"
	"time"
	"unicode"
)

func symbols() *[]string {
	symbols := []string{
		"enable",
		"public",
		"grade",
		"rocket",
		"cookie",
		"thunderstorm",
		"face",
		"skull",
		"home",
		"mode_cool",
		"bedroom_baby",
		"flatware",
		"single_bed",
		"sprinkler",
		"umbrella",
		"token",
		"skillet",
		"stadia_controller",
		"airwave",
		"floor_lamp",
		"close",
		"quiet_time",
		"heat",
		"tools_power_drill",
		"nest_eco_leaf",
		"air_freshener",
	}

	return &symbols
}

func LettersToRandomSymbols() map[string]string {
	letters := make(map[string]string)
	symbols := *symbols()
	randomizer := rand.New(rand.NewSource(time.Now().Unix()))
	randomIndexes := randomizer.Perm(len(symbols))
	index := 0
	for letter := 'a'; letter <= 'z'; letter++ {
		Letter := string(unicode.ToUpper(letter))
		letters[Letter] = symbols[randomIndexes[index]]
		index++
	}

	return letters
}
