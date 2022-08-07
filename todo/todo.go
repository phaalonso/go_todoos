package todo

import (
	"encoding/json"
	"io/ioutil"
)

type Item struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func SaveItems(fileName string, items []Item) error {
	b, err := json.Marshal(items)

	if err != nil {
		return err
	}

	// fmt.Println(string(b))

	err = ioutil.WriteFile(fileName, b, 0644)

	if err != nil {
		return err
	}

	return nil
}

func ReadItems(fileName string) ([]Item, error) {
	b, err := ioutil.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var items []Item

	if err := json.Unmarshal(b, &items); err != nil {
		return []Item{}, nil
	}

	return items, nil
}
