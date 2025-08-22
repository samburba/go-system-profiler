package displays

import (
	"encoding/json"
	"testing"
)

func TestDisplaysDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Fatalf("Failed to initialize DataType: %v", err)
	}

	// Test that DataType is not nil
	if DataType == nil {
		t.Fatal("DataType should not be nil")
	}

	// Test JSON marshaling
	jsonData, err := json.Marshal(DataType)
	if err != nil {
		t.Errorf("Failed to marshal DataType to JSON: %v", err)
	}

	// Test that JSON is not empty
	if len(jsonData) == 0 {
		t.Error("JSON data should not be empty")
	}

	// Test String() method
	str := DataType.String()
	if str == "" {
		t.Error("String() method should not return empty string")
	}

	// Verify JSON structure
	var parsed map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		t.Errorf("Failed to parse JSON: %v", err)
	}

	// Check for required top-level fields
	if _, exists := parsed["_name"]; !exists {
		t.Error("JSON should contain '_name' field")
	}
}

func TestDisplaysFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No displays data found")
	}

	// Test that we can access displays fields
	if name, exists := DataType["_name"]; !exists || name == "" {
		t.Error("Displays should have a name")
	}

	// Test that we can access other displays fields
	if vendor, exists := DataType["spdisplays_vendor"]; exists && vendor != "" {
		t.Logf("Display vendor: %v", vendor)
	}

	if model, exists := DataType["sppci_model"]; exists && model != "" {
		t.Logf("Display model: %v", model)
	}

	if cores, exists := DataType["sppci_cores"]; exists && cores != "" {
		t.Logf("Display cores: %v", cores)
	}
}
