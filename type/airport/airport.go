// Package airport provides data structures and initialization for system profiler's SPAirPortDataType.
package airport

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// WirelessNetwork represents a wireless network.
type WirelessNetwork struct {
	Name                    string `json:"_name"`
	SpairportNetworkChannel string `json:"spairport_network_channel,omitempty"`
	SpairportNetworkPhymode string `json:"spairport_network_phymode,omitempty"`
	SpairportNetworkType    string `json:"spairport_network_type,omitempty"`
	SpairportSecurityMode   string `json:"spairport_security_mode,omitempty"`
}

// AirportInterface represents an airport interface.
type AirportInterface struct {
	Name                                    string            `json:"_name"`
	SpairportAirportOtherLocalWirelessNetworks []WirelessNetwork `json:"spairport_airport_other_local_wireless_networks,omitempty"`
}

// SoftwareInformation represents airport software information.
type SoftwareInformation struct {
	SpairportCorewlanVersion      string `json:"spairport_corewlan_version,omitempty"`
	SpairportCorewlankitVersion   string `json:"spairport_corewlankit_version,omitempty"`
	SpairportDiagnosticsVersion   string `json:"spairport_diagnostics_version,omitempty"`
	SpairportExtraVersion         string `json:"spairport_extra_version,omitempty"`
	SpairportFamilyVersion        string `json:"spairport_family_version,omitempty"`
	SpairportProfilerVersion      string `json:"spairport_profiler_version,omitempty"`
	SpairportUtilityVersion       string `json:"spairport_utility_version,omitempty"`
}

// DataTypeItem represents the structure of SPAirPortDataType.
type DataTypeItem struct {
	SpairportAirportInterfaces []AirportInterface `json:"spairport_airport_interfaces,omitempty"`
	SpairportSoftwareInformation SoftwareInformation `json:"spairport_software_information,omitempty"`
}

// DataType holds the parsed system profiler data for SPAirPortDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPAirPortDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize airport data: %w", err)
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

