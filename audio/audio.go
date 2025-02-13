package audio

import (
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

const ThisDataType = "SPAudioDataType"

type DataTypeItem struct {
	Name                              string `json:"_name"`
	CoreaudioDefaultAudioInputDevice  string `json:"coreaudio_default_audio_input_device,omitempty"`
	CoreaudioDeviceInput              int    `json:"coreaudio_device_input,omitempty"`
	CoreaudioDeviceManufacturer       string `json:"coreaudio_device_manufacturer"`
	CoreaudioDeviceSrate              int    `json:"coreaudio_device_srate"`
	CoreaudioDeviceTransport          string `json:"coreaudio_device_transport"`
	CoreaudioInputSource              string `json:"coreaudio_input_source,omitempty"`
	Properties                        string `json:"_properties,omitempty"`
	CoreaudioDefaultAudioOutputDevice string `json:"coreaudio_default_audio_output_device,omitempty"`
	CoreaudioDefaultAudioSystemDevice string `json:"coreaudio_default_audio_system_device,omitempty"`
	CoreaudioDeviceOutput             int    `json:"coreaudio_device_output,omitempty"`
	CoreaudioOutputSource             string `json:"coreaudio_output_source,omitempty"`
}

var Data *common.Data = common.NewData(ThisDataType)
