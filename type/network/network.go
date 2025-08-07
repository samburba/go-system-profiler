// Package network provides data structures and initialization for system profiler's SPNetworkDataType.
package network

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// Ethernet represents ethernet configuration.
type Ethernet struct {
	MacAddress   string   `json:"MAC Address,omitempty"`
	MediaOptions []string `json:"MediaOptions,omitempty"`
	MediaSubType string   `json:"MediaSubType,omitempty"`
}

// IPv4 represents IPv4 configuration.
type IPv4 struct {
	ConfigMethod string `json:"ConfigMethod,omitempty"`
}

// IPv6 represents IPv6 configuration.
type IPv6 struct {
	ConfigMethod string `json:"ConfigMethod,omitempty"`
}

// Proxies represents proxy configuration.
type Proxies struct {
	ExceptionsList []string `json:"ExceptionsList,omitempty"`
	FTPPassive     string   `json:"FTPPassive,omitempty"`
}

// DataTypeItem represents the structure of SPNetworkDataType.
type DataTypeItem struct {
	Name                  string   `json:"_name"`
	Ethernet              Ethernet `json:"Ethernet,omitempty"`
	Hardware              string   `json:"hardware,omitempty"`
	Interface             string   `json:"interface,omitempty"`
	IPv4                  IPv4     `json:"IPv4,omitempty"`
	IPv6                  IPv6     `json:"IPv6,omitempty"`
	Proxies               Proxies  `json:"Proxies,omitempty"`
	SpnetworkServiceOrder int      `json:"spnetwork_service_order,omitempty"`
	Type                  string   `json:"type,omitempty"`
}

// DataType holds the parsed system profiler data for SPNetworkDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPNetworkDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize network data: %w", err)
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
