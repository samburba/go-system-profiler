package printerssoftware

import (
	"encoding/json"
	"testing"
)

func TestPrintersSoftwareDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	// Test that DataType is not nil
	if DataType == nil {
		t.Skip("Skipping: No printers software data available")
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

func TestPrintersSoftwareFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("Skipping: No printers software data found")
	}

	// Test that we can access printers software fields
	if len(DataType.Item) == 0 {
		t.Log("No printers software data found (this is normal if no printers software is available)")
		return
	}

	// Test that each printers software item has basic fields
	for i, item := range DataType.Item {
		if item.Name == "" {
			t.Errorf("Printers software item %d should have a name", i)
		}

		t.Logf("Printers software %d: %s", i, item.Name)
	}
}
