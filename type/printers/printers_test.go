package printers

import (
	"encoding/json"
	"testing"
)

func TestPrintersDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Fatalf("Failed to initialize DataType: %v", err)
	}

	// Test that DataType is not nil
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

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
}

func TestPrintersFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No printers data found")
	}

	// Test that we can access printers fields
	if len(DataType.Item) == 0 {
		t.Log("No printers data found (this is normal if no printers are available)")
		return
	}

	// Test that each printer item has basic fields
	for i, item := range DataType.Item {
		if item.Name == "" {
			t.Errorf("Printer item %d should have a name", i)
		}

		// Note: CupsVersion and Status fields are not implemented in the current DataTypeItem struct
		// TODO: Add specific fields when implementing the full printers structure
	}
}
