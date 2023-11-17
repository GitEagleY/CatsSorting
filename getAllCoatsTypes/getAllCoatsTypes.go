package main

import (
	"strings"

	"github.com/GitEagleY/CatsSorting/models"
	"github.com/GitEagleY/CatsSorting/processing"
)

func getUniqueCoatLengths(cats []models.Cat) []string {
	coatLengthsMap := make(map[string]bool)

	for _, cat := range cats {
		coat := strings.TrimSpace(cat.Coat)
		if coat != "" {
			coatLengthsMap[coat] = true
		}
	}

	var uniqueCoatLengths []string
	for length := range coatLengthsMap {
		uniqueCoatLengths = append(uniqueCoatLengths, length)
	}

	return uniqueCoatLengths
}

func main() {
	url := "https://catfact.ninja/breeds"

	//Getting data
	catsToWorkWith := processing.ParseData(url)

	coats := getUniqueCoatLengths(catsToWorkWith)

	processing.WriteToFile("coatsTypes", coats)

}
