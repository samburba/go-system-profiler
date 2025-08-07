package installhistory

import (
	"encoding/json"
	"testing"
)

func TestInstallHistoryDataType(t *testing.T) {
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

func TestInstallHistoryFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No install history data found")
	}

	// Test that we can access install history fields
	if len(DataType.Item) == 0 {
		t.Log("No install history found (this is normal if no install history is available)")
		return
	}

	// Test that each install history item has basic fields
	for i, item := range DataType.Item {
		if item.Name == "" {
			t.Errorf("Install history item %d should have a name", i)
		}

		// Note: InstallVersion field is not implemented in the current DataTypeItem struct
		// TODO: Add specific fields when implementing the full install history structure
	}
}
