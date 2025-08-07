// Package secureelement provides data structures and initialization for system profiler's SPSecureElementDataType.
package secureelement

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPSecureElementDataType.
type DataTypeItem struct {
	CtlFw              string `json:"ctl_fw,omitempty"`
	CtlHw              string `json:"ctl_hw,omitempty"`
	CtlInfo            string `json:"ctl_info,omitempty"`
	CtlMw              string `json:"ctl_mw,omitempty"`
	SeDevice           string `json:"se_device,omitempty"`
	SeFw               string `json:"se_fw,omitempty"`
	SeHw               string `json:"se_hw,omitempty"`
	SeID               string `json:"se_id,omitempty"`
	SeInRestrictedMode string `json:"se_in_restricted_mode,omitempty"`
	SeInfo             string `json:"se_info,omitempty"`
	SeOsID             string `json:"se_os_id,omitempty"`
	SeOsVersion        string `json:"se_os_version,omitempty"`
	SePlt              string `json:"se_plt,omitempty"`
	SeProdSigned       string `json:"se_prod_signed,omitempty"`
}

// DataType holds the parsed system profiler data for SPSecureElementDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPSecureElementDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize secure element data: %w", err)
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
