package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

// Add an item to the list
func (l *List) Add(task string) {
	*l = append(*l, item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	})
}

// Complete method marks a TODO item as completed by setting the Done field to true and the CompletedAt field to the current time.
func (l *List) Complete(index int) error {
	ls := *l
	if index <= 0 || index > len(ls) {
		return fmt.Errorf("Item %d out of range", index)
	}
	// Adjusting index for zero-based indexing
	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()
	return nil
}

// Delete method deletes a TODO item from the list by index.
func (l *List) Delete(index int) error {
	ls := *l
	if index <= 0 || index > len(ls) {
		return fmt.Errorf("Item %d out of range", index)
	}
	// Adjusting index for zero-based indexing
	*l = append(ls[:index-1], ls[index:]...)
	return nil
}

// Save method saves the list to a file.
func (l *List) Save(filename string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, js, 0644)
}

// Get method opens the provided file name, decodes the JSON data and parse it into a List.
func (l *List) Get(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	if len(file) == 0 {
		return nil
	}
	return json.Unmarshal(file, l)
}
