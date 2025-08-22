// Package memory provides data structures and initialization for system profiler's SPMemoryDataType.
package memory

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPMemoryDataType.
type DataTypeItem struct {
	Name string `json:"_name"`
}

// DataType holds the parsed system profiler data for SPMemoryDataType.
var DataType common.ObjectDataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewObjectData[DataTypeItem](common.SPMemoryDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize memory data: %w", err)
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

