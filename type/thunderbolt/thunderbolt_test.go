package thunderbolt

import (
	"encoding/json"
	"testing"
)

func TestThunderboltDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	// Test that DataType is not nil
	if DataType == nil {
		t.Skip("Skipping: No thunderbolt data available")
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

func TestThunderboltFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("Skipping: No thunderbolt data found")
	}

	// Test that we can access thunderbolt fields
	if len(DataType.Item) == 0 {
		t.Log("No thunderbolt data found (this is normal if no thunderbolt is available)")
		return
	}

	// Test that each thunderbolt item has basic fields
	for i, item := range DataType.Item {
		if item.Name == "" {
			t.Errorf("Thunderbolt item %d should have a name", i)
		}

		t.Logf("Thunderbolt device %d: %s", i, item.Name)
	}
}
