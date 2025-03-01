package common

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type DataType[T any] struct {
	Item []T   `json:"_items"`
	Name string `json:"_name"`
}

func (d *DataType[T]) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "Error converting Data to JSON string"
	}
	return string(jsonData)
}

func NewData[T any](spType SPDataType) (*DataType[T], error) {
	d, err := executeSPCommand[T](spType)
	if err != nil {
		return nil, err
	}
	return d, nil
}

func executeSPCommand[T any](spType SPDataType) (*DataType[T], error) {
	cmd := exec.Command("system_profiler", string(spType), "-json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}

	// Define the structure based on the data you expect
	var rawData map[string][]DataType[T]

	err = json.Unmarshal(output, &rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if items, exists := rawData[string(spType)]; exists && len(items) > 0 {
		return &items[0], nil // Return the first DataType item
	}

	return nil, fmt.Errorf("no data found for %s", spType)
}

