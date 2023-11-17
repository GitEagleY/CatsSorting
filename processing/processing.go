package processing

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/GitEagleY/CatsSorting/models"
)

func GetJsonData(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
func ParseData(url string) []models.Cat {

	jsonData, err := GetJsonData(url)
	if err != nil {
		log.Fatal("Error:", err)
		return nil
	}

	var catApiResponse models.CatApiResponse
	err = json.Unmarshal([]byte(jsonData), &catApiResponse)
	if err != nil {
		log.Fatal("Error:", err)
		return nil
	}

	var outputCats []models.Cat

	for _, catFromJson := range catApiResponse.Data {
		if err != nil {
			log.Fatal("Error:", err)
			return nil
		}
		outputCats = append(outputCats, catFromJson)
	}
	return outputCats

}
func WriteToFile(filename string, data []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, line := range data {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	return nil
}

func WriteJSON(filename string, data map[int]string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		return err
	}

	fmt.Printf("written to %s\n", filename)
	return nil
}
