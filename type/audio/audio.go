// Package audio provides data structures and initialization for system profiler's SPAudioDataType.
package audio

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPAudioDataType.
type DataTypeItem struct {
	Name                              string `json:"_name"`
	Properties                        string `json:"_properties,omitempty"`
	CoreaudioDefaultAudioInputDevice  string `json:"coreaudio_default_audio_input_device,omitempty"`
	CoreaudioDeviceInput              int    `json:"coreaudio_device_input,omitempty"`
	CoreaudioDeviceManufacturer       string `json:"coreaudio_device_manufacturer,omitempty"`
	CoreaudioDeviceSrate              int    `json:"coreaudio_device_srate,omitempty"`
	CoreaudioDeviceTransport          string `json:"coreaudio_device_transport,omitempty"`
	CoreaudioInputSource              string `json:"coreaudio_input_source,omitempty"`
	CoreaudioDefaultAudioOutputDevice string `json:"coreaudio_default_audio_output_device,omitempty"`
	CoreaudioDefaultAudioSystemDevice string `json:"coreaudio_default_audio_system_device,omitempty"`
	CoreaudioDeviceOutput             int    `json:"coreaudio_device_output,omitempty"`
	CoreaudioOutputSource             string `json:"coreaudio_output_source,omitempty"`
}

// DataType holds the parsed system profiler data for SPAudioDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPAudioDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize audio data: %w", err)
			return
		}
		DataType = data
	})
	return initErr
}

// GetDataType returns the DataType, initializing it if necessary
func GetDataType() (*common.DataType[DataTypeItem], error) {
	if err := Initialize(); err != nil {
		return nil, err
	}
	return DataType, nil
}
