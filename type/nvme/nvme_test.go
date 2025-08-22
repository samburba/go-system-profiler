package nvme

import (
	"encoding/json"
	"testing"
)

func TestNVMeDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	// Test that DataType is not nil
	if DataType == nil {
		t.Skip("Skipping: No NVMe data available")
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

func TestNVMeFields(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Skipping: Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("Skipping: No NVMe data found")
	}

	// Test that we can access NVMe fields
	if len(DataType.Item) == 0 {
		t.Log("No NVMe data found (this is normal if no NVMe devices are available)")
		return
	}

	// Test that each NVMe item has basic fields
	for i, item := range DataType.Item {
		if item.Name == "" {
			t.Errorf("NVMe item %d should have a name", i)
		}

		// Test NVMe devices if available
		if len(item.Items) > 0 {
			for j, device := range item.Items {
				if device.Name == "" {
					t.Errorf("NVMe device %d in item %d should have a name", j, i)
				}

				if device.SizeInBytes <= 0 {
					t.Errorf("NVMe device %d in item %d should have a size", j, i)
				}

				t.Logf("NVMe Device: %s (%s)", device.Name, device.Size)

				// Test volumes if available
				if len(device.Volumes) > 0 {
					for k, volume := range device.Volumes {
						t.Logf("  Volume %d: %s (%s)", k, volume.Name, volume.Size)
					}
				}
			}
		}
	}
}
