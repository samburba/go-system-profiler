// Package displays provides data structures and initialization for system profiler's SPDisplaysDataType.
package displays

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPDisplaysDataType.
type DataTypeItem struct {
	Name                          string `json:"_name"`
	SpdisplaysMtlgpufamilysupport string `json:"spdisplays_mtlgpufamilysupport,omitempty"`
	SpdisplaysNdrvs               string `json:"spdisplays_ndrvs,omitempty"`
	SpdisplaysVendor              string `json:"spdisplays_vendor,omitempty"`
	SppciBus                      string `json:"sppci_bus,omitempty"`
	SppciCores                    string `json:"sppci_cores,omitempty"`
	SppciDeviceType               string `json:"sppci_device_type,omitempty"`
	SppciModel                    string `json:"sppci_model,omitempty"`
}

// DataType holds the parsed system profiler data for SPDisplaysDataType.
var DataType common.ObjectDataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewObjectData[DataTypeItem](common.SPDisplaysDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize displays data: %w", err)
			return
		}
		DataType = data
	})
	return initErr
}

// GetDataType returns the DataType, initializing it if necessary
func GetDataType() (common.ObjectDataType[DataTypeItem], error) {
	if err := Initialize(); err != nil {
		return nil, err
	}
	return DataType, nil
}
