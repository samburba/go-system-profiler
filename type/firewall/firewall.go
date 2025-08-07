// Package firewall provides data structures and initialization for system profiler's SPFirewallDataType.
package firewall

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// Applications represents firewall applications.
type Applications struct {
	ComAppleCupsd              string `json:"com.apple.cupsd,omitempty"`
	ComAppleDtXcodeSelectToolShim string `json:"com.apple.dt.xcode_select.tool-shim,omitempty"`
	ComAppleRemoted            string `json:"com.apple.remoted,omitempty"`
	ComAppleRuby               string `json:"com.apple.ruby,omitempty"`
	ComAppleSharingd           string `json:"com.apple.sharingd,omitempty"`
	ComAppleSmbd               string `json:"com.apple.smbd,omitempty"`
	ComAppleSshdKeygenWrapper  string `json:"com.apple.sshd-keygen-wrapper,omitempty"`
}

// DataTypeItem represents the structure of SPFirewallDataType.
type DataTypeItem struct {
	Name                      string       `json:"_name"`
	SpfirewallApplications    Applications `json:"spfirewall_applications,omitempty"`
	SpfirewallGlobalstate     string       `json:"spfirewall_globalstate,omitempty"`
	SpfirewallLoggingenabled  string       `json:"spfirewall_loggingenabled,omitempty"`
	SpfirewallStealthenabled  string       `json:"spfirewall_stealthenabled,omitempty"`
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
