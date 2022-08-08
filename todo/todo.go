package todo

import (
	"encoding/json"
	"fmt"
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

func ListItems(items []Item) {
	for i, x := range items {
		checked := " "
		if x.Done {
			checked = "✓"
		}
		fmt.Printf("%s %d - %s\n", checked, i+1, x.Text)
	}
}
