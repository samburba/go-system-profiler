// Package software provides data structures and initialization for system profiler's SPSoftwareDataType.
package software

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPSoftwareDataType.
type DataTypeItem struct {
	Name             string `json:"_name"`
	BootMode         string `json:"boot_mode,omitempty"`
	BootVolume       string `json:"boot_volume,omitempty"`
	KernelVersion    string `json:"kernel_version,omitempty"`
	LocalHostName    string `json:"local_host_name,omitempty"`
	OsVersion        string `json:"os_version,omitempty"`
	SecureVM         string `json:"secure_vm,omitempty"`
	SystemIntegrity  string `json:"system_integrity,omitempty"`
	Uptime           string `json:"uptime,omitempty"`
	UserName         string `json:"user_name,omitempty"`
}

// DataType holds the parsed system profiler data for SPSoftwareDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPSoftwareDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize software data: %w", err)
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
