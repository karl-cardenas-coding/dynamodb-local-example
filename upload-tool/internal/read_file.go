package internal

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Get table items from JSON file
func GetObjectsFromFile(dataPath string) []Model {
	raw, err := ioutil.ReadFile(dataPath)
	if err != nil {
		log.Fatal("Error reading mock data file", err)
	}

	var items []Model
	err = json.Unmarshal(raw, &items)
	if err != nil {
		log.Fatal("Unmarshall Error:", err)
	}
	return items
}
