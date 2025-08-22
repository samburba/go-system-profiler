// Package applications provides data structures and initialization for system profiler's SPApplicationsDataType.
package applications

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// DataTypeItem represents the structure of SPApplicationsDataType.
type DataTypeItem struct {
	Name         string   `json:"_name"`
	ArchKind     string   `json:"arch_kind,omitempty"`
	LastModified string   `json:"lastModified,omitempty"`
	ObtainedFrom string   `json:"obtained_from,omitempty"`
	Path         string   `json:"path,omitempty"`
	SignedBy     []string `json:"signed_by,omitempty"`
	Version      string   `json:"version,omitempty"`
}

// DataType holds the parsed system profiler data for SPApplicationsDataType.
var DataType common.DirectDataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewDirectData[DataTypeItem](common.SPApplicationsDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize applications data: %w", err)
			return
		}
		DataType = data
	})
	return initErr
}

// GetDataType returns the DataType, initializing it if necessary
func GetDataType() (common.DirectDataType[DataTypeItem], error) {
	if err := Initialize(); err != nil {
		return nil, err
	}
	return DataType, nil
}
