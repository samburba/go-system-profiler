package memory

import (
	"encoding/json"
	"testing"
)

func TestMemoryDataType(t *testing.T) {
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

	// Check for memory-specific fields
	if _, exists := parsed["SPMemoryDataType"]; !exists {
		t.Error("JSON should contain 'SPMemoryDataType' field")
	}
}

func TestMemoryFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No memory data found")
	}

	// Test that we can access memory fields
	if dimmType, exists := DataType["dimm_type"]; exists && dimmType != "" {
		t.Logf("Memory DIMM type: %v", dimmType)
	}

	if manufacturer, exists := DataType["dimm_manufacturer"]; exists && manufacturer != "" {
		t.Logf("Memory manufacturer: %v", manufacturer)
	}

	// Log memory information
	t.Logf("Memory data structure test passed")
}
