// Package hardware provides data structures and initialization for system profiler's SPHardwareDataType.
package hardware

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPHardwareDataType.
type DataTypeItem struct {
	Name                  string `json:"_name"`
	ActivationLockStatus  string `json:"activation_lock_status,omitempty"`
	BootRomVersion        string `json:"boot_rom_version,omitempty"`
	ChipType              string `json:"chip_type,omitempty"`
	MachineModel          string `json:"machine_model,omitempty"`
	MachineName           string `json:"machine_name,omitempty"`
	ModelNumber           string `json:"model_number,omitempty"`
	NumberProcessors      string `json:"number_processors,omitempty"`
	OsLoaderVersion       string `json:"os_loader_version,omitempty"`
	PhysicalMemory        string `json:"physical_memory,omitempty"`
	PlatformUUID          string `json:"platform_UUID,omitempty"`
	ProvisioningUDID      string `json:"provisioning_UDID,omitempty"`
	SerialNumber          string `json:"serial_number,omitempty"`
}

// DataType holds the parsed system profiler data for SPHardwareDataType.
var DataType common.ObjectDataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewObjectData[DataTypeItem](common.SPHardwareDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize hardware data: %w", err)
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
