// Package displays provides data structures and initialization for system profiler's SPDisplaysDataType.
package displays

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DisplayDriver represents a display driver/device.
type DisplayDriver struct {
	Name                        string `json:"_name"`
	SpdisplaysDisplayProductID  string `json:"_spdisplays_display-product-id,omitempty"`
	SpdisplaysDisplaySerialNumber string `json:"_spdisplays_display-serial-number,omitempty"`
	SpdisplaysDisplayVendorID   string `json:"_spdisplays_display-vendor-id,omitempty"`
	SpdisplaysDisplayWeek       string `json:"_spdisplays_display-week,omitempty"`
	SpdisplaysDisplayYear       string `json:"_spdisplays_display-year,omitempty"`
	SpdisplaysDisplayID         string `json:"_spdisplays_displayID,omitempty"`
	SpdisplaysPixels            string `json:"_spdisplays_pixels,omitempty"`
	SpdisplaysResolution        string `json:"_spdisplays_resolution,omitempty"`
	SpdisplaysAmbientBrightness string `json:"spdisplays_ambient_brightness,omitempty"`
	SpdisplaysConnectionType    string `json:"spdisplays_connection_type,omitempty"`
	SpdisplaysDisplayType       string `json:"spdisplays_display_type,omitempty"`
	SpdisplaysMain              string `json:"spdisplays_main,omitempty"`
	SpdisplaysMirror            string `json:"spdisplays_mirror,omitempty"`
	SpdisplaysOnline            string `json:"spdisplays_online,omitempty"`
	SpdisplaysPixelResolution   string `json:"spdisplays_pixelresolution,omitempty"`
}

// DataTypeItem represents the structure of SPDisplaysDataType.
type DataTypeItem struct {
	Name                    string           `json:"_name"`
	SpdisplaysMTLGPUFamilySupport string `json:"spdisplays_mtlgpufamilysupport,omitempty"`
	SpdisplaysNDrvs        []DisplayDriver  `json:"spdisplays_ndrvs,omitempty"`
	SpdisplaysVendor       string           `json:"spdisplays_vendor,omitempty"`
	SppciBus               string           `json:"sppci_bus,omitempty"`
	SppciCores             string           `json:"sppci_cores,omitempty"`
	SppciDeviceType        string           `json:"sppci_device_type,omitempty"`
	SppciModel             string           `json:"sppci_model,omitempty"`
}

// DataType holds the parsed system profiler data for SPDisplaysDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPDisplaysDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize displays data: %w", err)
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

