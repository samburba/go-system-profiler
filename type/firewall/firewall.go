// Package firewall provides data structures and initialization for system profiler's SPFirewallDataType.
package firewall

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPFirewallDataType.
type DataTypeItem struct {
	Name                     string            `json:"_name"`
	SpfirewallApplications   map[string]string `json:"spfirewall_applications,omitempty"`
	SpfirewallGlobalState    string            `json:"spfirewall_globalstate,omitempty"`
	SpfirewallLoggingEnabled string            `json:"spfirewall_loggingenabled,omitempty"`
	SpfirewallStealthEnabled string            `json:"spfirewall_stealthenabled,omitempty"`
}

// DataType holds the parsed system profiler data for SPFirewallDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPFirewallDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize firewall data: %w", err)
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
