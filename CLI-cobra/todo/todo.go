package todo

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Item struct {
	Text     string
	Priority int
	position int
	Done     bool
}

type ByPriority []Item

func (s ByPriority) Len() int      { return len(s) }
func (s ByPriority) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s ByPriority) Less(i, j int) bool {
	if s[i].Done != s[j].Done {
		return s[i].Done
	}
	if s[i].Priority != s[j].Priority {
		return s[i].position < s[j].position
	}
	return s[i].Priority < s[j].Priority
}

func SaveItems(filename string, items []Item) error {
	fmt.Println("Writing to file:", filename) // <- debug print
	data, err := json.Marshal(items)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func ReadItems(filename string) ([]Item, error) {
	// Implementation to read items from a file
	data, err := os.ReadFile(filename)
	if err != nil {
		return []Item{}, err
	}
	var items []Item
	if err := json.Unmarshal(data, &items); err != nil {
		return []Item{}, err
	}
	for i := range items {
		items[i].position = i + 1
	}
	return items, nil
}

func (i *Item) SetPtiority(pri int) {
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

// Prettie priority print
func (i Item) PrettyP() string {
	if i.Priority == 1 {
		return "High"
	}
	if i.Priority == 3 {
		return "Low"
	}
	return "Medium"
}

func (i Item) Label() string {
	return strconv.Itoa(i.position) + ". "
}

func (i *Item) PrettyDone() string {
	if i.Done {
		return "[x] "
	}
	return "[ ] "
}
