package todo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sort"
)

type Item struct {
	Text     string `json:"text"`
	Priority int
	Done     bool `json:"done"`
	position int
}

func (i *Item) SetPriority(pri int) {
	switch pri {
	case 1:
		i.Priority = 1
	case 2:
		i.Priority = 2
	case 3:
		i.Priority = 3
	default:
		i.Priority = 2
	}
}

func (i *Item) PrettyP() string {
	if i.Priority == 1 {
		return "(1)"
	}
	if i.Priority == 3 {
		return "(3)"
	}
	return " "
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "X"
	}
	return ""
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

	for i := range items {
		items[i].position = i + 1
	}

	return items, nil
}

func ListItems(items []Item) {
	sort.Slice(items, func(i, j int) bool {
		return items[i].Priority > items[j].Priority
	})

	for i, x := range items {
		checked := " "
		if x.Done {
			checked = "✓"
		}
		fmt.Printf("%s %d - %s\n", checked, i+1, x.Text)
	}
}

// ByPri implements sort.Interface for []Item based on
// the Priority & position field

type ByPri []Item

func (s ByPri) Len() int      { return len(s) }
func (s ByPri) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPri) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}

	// If the items priority is the same, an item will be considered
	// than the other by its position number
	if s[i].Priority == s[j].Priority {
		return s[i].position < s[j].position
	}

	return s[i].Priority < s[j].Priority
}
