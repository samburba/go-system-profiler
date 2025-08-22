package common

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

// DataType represents the structure for items-based data (like Audio)
type DataType[T any] struct {
	Item []T    `json:"_items"`
	Name string `json:"_name"`
}

// DirectDataType represents the structure for direct array data (like Applications)
type DirectDataType[T any] []T

// ObjectDataType represents the structure for object data (like Network, Bluetooth)
type ObjectDataType[T any] map[string]interface{}

func (d *DataType[T]) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "Error converting Data to JSON string"
	}
	return string(jsonData)
}

func (d DirectDataType[T]) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "Error converting DirectData to JSON string"
	}
	return string(jsonData)
}

func (d ObjectDataType[T]) String() string {
	jsonData, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return "Error converting ObjectData to JSON string"
	}
	return string(jsonData)
}

// NewData creates a DataType for items-based structures
func NewData[T any](spType SPDataType) (*DataType[T], error) {
	d, err := executeSPCommand[T](spType)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// NewDirectData creates a DirectDataType for direct array structures
func NewDirectData[T any](spType SPDataType) (DirectDataType[T], error) {
	d, err := executeDirectSPCommand[T](spType)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// NewObjectData creates an ObjectDataType for object structures
func NewObjectData[T any](spType SPDataType) (ObjectDataType[T], error) {
	d, err := executeObjectSPCommand[T](spType)
	if err != nil {
		return nil, err
	}
	return d, nil
}

// executeSPCommand handles items-based structures (like Audio)
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

// executeDirectSPCommand handles direct array structures (like Applications)
func executeDirectSPCommand[T any](spType SPDataType) (DirectDataType[T], error) {
	cmd := exec.Command("system_profiler", string(spType), "-json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}

	var rawData map[string]DirectDataType[T]

	err = json.Unmarshal(output, &rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if items, exists := rawData[string(spType)]; exists {
		return items, nil
	}

	return nil, fmt.Errorf("no data found for %s", spType)
}

// executeObjectSPCommand handles object structures (like Network, Bluetooth)
func executeObjectSPCommand[T any](spType SPDataType) (ObjectDataType[T], error) {
	cmd := exec.Command("system_profiler", string(spType), "-json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}

	var rawData map[string][]ObjectDataType[T]

	err = json.Unmarshal(output, &rawData)
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	if items, exists := rawData[string(spType)]; exists && len(items) > 0 {
		return items[0], nil
	}

	return nil, fmt.Errorf("no data found for %s", spType)
}
