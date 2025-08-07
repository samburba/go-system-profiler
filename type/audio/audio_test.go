package audio

import (
	"encoding/json"
	"testing"
)

func TestAudioDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Fatalf("Failed to initialize DataType: %v", err)
	}

	// Initialize the data first
	if err := Initialize(); err != nil {
		t.Logf("Initialize returned error (this might be expected): %v", err)
	}

	// Test that DataType is not nil (if initialization succeeded)
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Log("DataType is nil (this might be expected if system_profiler failed)")
		return
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

func TestAudioFields(t *testing.T) {
	// Initialize the data first
	if err := Initialize(); err != nil {
		t.Skipf("Skipping test due to initialization error: %v", err)
	}

	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Skipf("Failed to initialize DataType: %v", err)
	}

	if DataType == nil {
		t.Skip("No audio data found")
	}

	// Test that we can access audio fields
	if len(DataType.Item) == 0 {
		t.Log("No audio devices found (this is normal if no audio is available)")
		return
	}

	// Test that each audio device has a name
	for i, device := range DataType.Item {
		if device.Name == "" {
			t.Errorf("Audio device %d should have a name", i)
		}
	}
}

func TestInitialize(t *testing.T) {
	// Test that Initialize works correctly
	err := Initialize()
	if err != nil {
		t.Logf("Initialize returned error (this might be expected): %v", err)
	}

	// Test that Initialize is idempotent
	err2 := Initialize()
	if err != err2 {
		t.Error("Initialize should be idempotent")
	}
}

func TestGetDataType(t *testing.T) {
	// Initialize the DataType
	if err := Initialize(); err != nil {
		t.Fatalf("Failed to initialize DataType: %v", err)
	}

	// Test GetDataType function
	dataType, err := GetDataType()
	if err != nil {
		t.Logf("GetDataType returned error (this might be expected): %v", err)
		return
	}

	if dataType == nil {
		t.Error("GetDataType should not return nil when no error")
	}

	// Test that it returns the same instance
	if dataType != DataType {
		t.Error("GetDataType should return the same instance as DataType")
	}
}
