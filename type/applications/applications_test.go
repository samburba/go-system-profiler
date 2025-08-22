package applications

import (
	"encoding/json"
	"testing"
)

func TestApplicationsDataType(t *testing.T) {
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
	var parsed []map[string]interface{}
	err = json.Unmarshal(jsonData, &parsed)
	if err != nil {
		t.Errorf("Failed to parse JSON: %v", err)
	}
}

func TestApplicationsFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No applications data found")
	}

	// Test that we can access applications fields
	if len(DataType) == 0 {
		t.Log("No applications data found (this is normal if no applications are available)")
		return
	}

	// Test that each applications item has basic fields
	for i, item := range DataType {
		if item.Name == "" {
			t.Errorf("Applications item %d should have a name", i)
		}

		t.Logf("Applications item %d: %s", i, item.Name)
	}
}
