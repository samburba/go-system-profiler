package common

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type Data struct {
	DataType map[string]any `json:"-"`
}

func (d *Data) String() string {
	jsonData, err := json.MarshalIndent(d.DataType, "", "  ")
	if err != nil {
		return "Error converting Data to JSON string"
	}
	return string(jsonData)
}

func NewData(sp_type string) *Data {
	d, err := executeSPCommand(sp_type)
	if err != nil {
		panic(err)
	}
	return d
}

func executeSPCommand(sp_type string) (*Data, error) {
	cmd := exec.Command("system_profiler", sp_type, "-json")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to execute command: %w", err)
	}
	var data Data
	err = json.Unmarshal(output, &data.DataType) // unmarshal into DataType directly
	if err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %w", err)
	}

	return &data, nil
}
