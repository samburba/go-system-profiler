// Package bluetooth provides data structures and initialization for system profiler's SPBluetoothDataType.
package bluetooth

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// ControllerProperties represents Bluetooth controller properties.
type ControllerProperties struct {
	ControllerAddress           string `json:"controller_address,omitempty"`
	ControllerChipset           string `json:"controller_chipset,omitempty"`
	ControllerDiscoverable      string `json:"controller_discoverable,omitempty"`
	ControllerFirmwareVersion   string `json:"controller_firmwareVersion,omitempty"`
	ControllerProductID         string `json:"controller_productID,omitempty"`
	ControllerState             string `json:"controller_state,omitempty"`
	ControllerSupportedServices string `json:"controller_supportedServices,omitempty"`
	ControllerTransport         string `json:"controller_transport,omitempty"`
	ControllerVendorID          string `json:"controller_vendorID,omitempty"`
}

// DeviceNotConnected represents a Bluetooth device that is not currently connected.
type DeviceNotConnected struct {
	DeviceAddress   string `json:"device_address,omitempty"`
	DeviceMinorType string `json:"device_minorType,omitempty"`
}

// DataTypeItem represents the structure of SPBluetoothDataType.
type DataTypeItem struct {
	ControllerProperties ControllerProperties            `json:"controller_properties,omitempty"`
	DeviceNotConnected   []map[string]DeviceNotConnected `json:"device_not_connected,omitempty"`
}

// DataType holds the parsed system profiler data for SPBluetoothDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPBluetoothDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize bluetooth data: %w", err)
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
