package main

import (
	"fmt"
	"sort"

	"github.com/GitEagleY/CatsSorting/models"
	"github.com/GitEagleY/CatsSorting/processing"
)

func main() {
	url := "https://catfact.ninja/breeds"

	coatTypes := map[string]int{
		"Hairless/Furry down": 1,
		"Short":               2,
		"Rex":                 3,
		"Medium":              4,
		"Semi-long":           5,
		"Short/Long":          6,
		"Semi Long":           7,
		"Long":                8,
	}

	// Getting data
	catsToWorkWith := processing.ParseData(url)
	fmt.Println(catsToWorkWith[0].Coat)

	// Output list of cat breeds
	for i, v := range catsToWorkWith {
		fmt.Printf("cat of number %d: has %s breed\n", i, v.Breed)
	}

	//grouping
	groupedCats := groupCatsByCountry(catsToWorkWith)

	//sorting
	sortedCats := sortByCoatLength(groupedCats, coatTypes)

	//processing
	processing.WriteJSON("out.json", sortedCats)

}

func groupCatsByCountry(cats []models.Cat) map[string][]models.Cat {
	groupedCats := make(map[string][]models.Cat)

	for _, cat := range cats {
		groupedCats[cat.Origin] = append(groupedCats[cat.Origin], cat)
	}
	return groupedCats
}

func sortByCoatLength(cats map[string][]models.Cat, coatTypes map[string]int) map[int]string {
	sortedCats := make(map[int]string)

	for _, group := range cats {

		sort.Slice(group, func(i, j int) bool {
			return coatTypes[group[i].Coat] < coatTypes[group[j].Coat]
		})

		for i, cat := range group {
			sortedCats[i+1] = cat.Breed
		}
	}

	return sortedCats
}
