// Package fonts provides data structures and initialization for system profiler's SPFontsDataType.
package fonts

import (
	"fmt"
	"sync"

	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// Typeface represents a font typeface.
type Typeface struct {
	Name          string `json:"_name"`
	CopyProtected string `json:"copy_protected,omitempty"`
	Copyright     string `json:"copyright,omitempty"`
	Designer      string `json:"designer,omitempty"`
	Duplicate     string `json:"duplicate,omitempty"`
	Embeddable    string `json:"embeddable,omitempty"`
	Enabled       string `json:"enabled,omitempty"`
	Family        string `json:"family,omitempty"`
	Fullname      string `json:"fullname,omitempty"`
	Outline       string `json:"outline,omitempty"`
	Style         string `json:"style,omitempty"`
	Unique        string `json:"unique,omitempty"`
	Valid         string `json:"valid,omitempty"`
}

// DataTypeItem represents the structure of SPFontsDataType.
type DataTypeItem struct {
	Name      string     `json:"_name"`
	Enabled   string     `json:"enabled,omitempty"`
	Path      string     `json:"path,omitempty"`
	Type      string     `json:"type,omitempty"`
	Typefaces []Typeface `json:"typefaces,omitempty"`
}

// DataType holds the parsed system profiler data for SPFontsDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPFontsDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize fonts data: %w", err)
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
