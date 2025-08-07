// Package storage provides data structures and initialization for system profiler's SPStorageDataType.
package storage

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// PhysicalDrive represents physical drive information.
type PhysicalDrive struct {
	DeviceName         string `json:"device_name,omitempty"`
	IsInternalDisk     string `json:"is_internal_disk,omitempty"`
	MediaName          string `json:"media_name,omitempty"`
	MediumType         string `json:"medium_type,omitempty"`
	PartitionMapType   string `json:"partition_map_type,omitempty"`
	Protocol           string `json:"protocol,omitempty"`
	SmartStatus        string `json:"smart_status,omitempty"`
}

// DataTypeItem represents the structure of SPStorageDataType.
type DataTypeItem struct {
	Name              string        `json:"_name"`
	BsdName           string        `json:"bsd_name,omitempty"`
	FileSystem        string        `json:"file_system,omitempty"`
	FreeSpaceInBytes  int64         `json:"free_space_in_bytes,omitempty"`
	IgnoreOwnership   string        `json:"ignore_ownership,omitempty"`
	MountPoint        string        `json:"mount_point,omitempty"`
	PhysicalDrive     PhysicalDrive `json:"physical_drive,omitempty"`
	SizeInBytes       int64         `json:"size_in_bytes,omitempty"`
	VolumeUUID        string        `json:"volume_uuid,omitempty"`
	Writable          string        `json:"writable,omitempty"`
}

// DataType holds the parsed system profiler data for SPStorageDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPStorageDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize storage data: %w", err)
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

