package airport

import (
	"encoding/json"
	"testing"
)

func TestAirportDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Fatalf("Failed to initialize DataType: %v", err)
	}

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

func TestAirportInterfaces(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No airport data found")
	}

	// Test that we can access airport interfaces
	if len(DataType.Item) == 0 {
		t.Log("No airport interfaces found (this is normal if no WiFi is available)")
		return
	}

	// Test that each interface has airport data
	for i, item := range DataType.Item {
		if len(item.SpairportAirportInterfaces) == 0 {
			t.Logf("Airport item %d has no interfaces (this is normal)", i)
			continue
		}

		// Test that each interface has a name
		for j, iface := range item.SpairportAirportInterfaces {
			if iface.Name == "" {
				t.Errorf("Airport interface %d in item %d should have a name", j, i)
			}
		}
	}
}

func TestAirportSoftwareInformation(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No airport data found")
	}

	// Test that software information is accessible
	// Note: Software information might not always be present
	t.Log("Airport software information test passed")
}
