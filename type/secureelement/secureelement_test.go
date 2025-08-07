package secureelement

import (
	"encoding/json"
	"testing"
)

func TestSecureElementDataType(t *testing.T) {
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

func TestSecureElementFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No secure element data found")
	}

	// Test that we can access secure element fields
	if len(DataType.Item) == 0 {
		t.Log("No secure element data found (this is normal if no secure element is available)")
		return
	}

	// Test that each secure element item has basic fields
	for i, item := range DataType.Item {
		if item.SeDevice == "" {
			t.Errorf("Secure element item %d should have a device ID", i)
		}

		if item.SeID == "" {
			t.Errorf("Secure element item %d should have an ID", i)
		}
	}
}
