// Package thunderbolt provides data structures and initialization for system profiler's SPThunderboltDataType.
package thunderbolt

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// ReceptacleTag represents Thunderbolt receptacle information.
type ReceptacleTag struct {
	CurrentSpeedKey     string `json:"current_speed_key,omitempty"`
	LinkStatusKey       string `json:"link_status_key,omitempty"`
	ReceptacleIDKey     string `json:"receptacle_id_key,omitempty"`
	ReceptacleStatusKey string `json:"receptacle_status_key,omitempty"`
}

// DataTypeItem represents the structure of SPThunderboltDataType.
type DataTypeItem struct {
	Name              string        `json:"_name"`
	DeviceNameKey     string        `json:"device_name_key,omitempty"`
	DomainUUIDKey     string        `json:"domain_uuid_key,omitempty"`
	Receptacle1Tag    ReceptacleTag `json:"receptacle_1_tag,omitempty"`
	RouteStringKey    string        `json:"route_string_key,omitempty"`
	SwitchUIDKey      string        `json:"switch_uid_key,omitempty"`
	VendorNameKey     string        `json:"vendor_name_key,omitempty"`
}

// DataType holds the parsed system profiler data for SPThunderboltDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPThunderboltDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize thunderbolt data: %w", err)
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

