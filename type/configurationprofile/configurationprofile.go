// Package configurationprofile provides data structures and initialization for system profiler's SPConfigurationProfileDataType.
package configurationprofile

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPConfigurationProfileDataType.
type DataTypeItem struct {
	Name string `json:"_name"`
	// TODO: Add specific fields based on system_profiler output
}

// DataType holds the parsed system profiler data for SPConfigurationProfileDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPConfigurationProfileDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize configurationprofile data: %w", err)
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
