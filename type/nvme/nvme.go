// Package nvme provides data structures and initialization for system profiler's SPNVMeDataType.
package nvme

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// Volume represents an NVMe volume.
type Volume struct {
	Name          string `json:"_name"`
	BsdName       string `json:"bsd_name,omitempty"`
	IOContent     string `json:"iocontent,omitempty"`
	Size          string `json:"size,omitempty"`
	SizeInBytes   int64  `json:"size_in_bytes,omitempty"`
}

// NVMeDevice represents an NVMe device.
type NVMeDevice struct {
	Name                string    `json:"_name"`
	BsdName             string    `json:"bsd_name,omitempty"`
	DetachableDrive     string    `json:"detachable_drive,omitempty"`
	DeviceModel         string    `json:"device_model,omitempty"`
	DeviceRevision      string    `json:"device_revision,omitempty"`
	DeviceSerial        string    `json:"device_serial,omitempty"`
	PartitionMapType    string    `json:"partition_map_type,omitempty"`
	RemovableMedia      string    `json:"removable_media,omitempty"`
	Size                string    `json:"size,omitempty"`
	SizeInBytes         int64     `json:"size_in_bytes,omitempty"`
	SmartStatus         string    `json:"smart_status,omitempty"`
	SpnvmeTrimSupport   string    `json:"spnvme_trim_support,omitempty"`
	Volumes             []Volume  `json:"volumes,omitempty"`
}

// DataTypeItem represents the structure of SPNVMeDataType.
type DataTypeItem struct {
	Name  string        `json:"_name"`
	Items []NVMeDevice  `json:"_items,omitempty"`
}

// DataType holds the parsed system profiler data for SPNVMeDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPNVMeDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize nvme data: %w", err)
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

