// Package power provides data structures and initialization for system profiler's SPPowerDataType.
package power

import (
	"fmt"
	"sync"
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

// BatteryChargeInfo represents battery charge information.
type BatteryChargeInfo struct {
	SppowerBatteryAtWarnLevel    string `json:"sppower_battery_at_warn_level,omitempty"`
	SppowerBatteryFullyCharged   string `json:"sppower_battery_fully_charged,omitempty"`
	SppowerBatteryIsCharging     string `json:"sppower_battery_is_charging,omitempty"`
	SppowerBatteryStateOfCharge  int    `json:"sppower_battery_state_of_charge,omitempty"`
}

// BatteryHealthInfo represents battery health information.
type BatteryHealthInfo struct {
	SppowerBatteryCycleCount           int    `json:"sppower_battery_cycle_count,omitempty"`
	SppowerBatteryHealth               string `json:"sppower_battery_health,omitempty"`
	SppowerBatteryHealthMaximumCapacity string `json:"sppower_battery_health_maximum_capacity,omitempty"`
}

// BatteryModelInfo represents battery model information.
type BatteryModelInfo struct {
	PackLotCode                    string `json:"Pack Lot Code,omitempty"`
	PCBLotCode                     string `json:"PCB Lot Code,omitempty"`
	SppowerBatteryCellRevision     string `json:"sppower_battery_cell_revision,omitempty"`
	SppowerBatteryDeviceName       string `json:"sppower_battery_device_name,omitempty"`
	SppowerBatteryFirmwareVersion  string `json:"sppower_battery_firmware_version,omitempty"`
	SppowerBatteryHardwareRevision string `json:"sppower_battery_hardware_revision,omitempty"`
	SppowerBatterySerialNumber     string `json:"sppower_battery_serial_number,omitempty"`
}

// PowerSettings represents power settings for AC or Battery.
type PowerSettings struct {
	CurrentPowerSource                    string `json:"Current Power Source,omitempty"`
	DiskSleepTimer                        int    `json:"Disk Sleep Timer,omitempty"`
	DisplaySleepTimer                     int    `json:"Display Sleep Timer,omitempty"`
	HibernateMode                         int    `json:"Hibernate Mode,omitempty"`
	LowPowerMode                          string `json:"LowPowerMode,omitempty"`
	PrioritizeNetworkReachabilityOverSleep string `json:"PrioritizeNetworkReachabilityOverSleep,omitempty"`
	ReduceBrightness                      string `json:"ReduceBrightness,omitempty"`
	SleepOnPowerButton                    string `json:"Sleep On Power Button,omitempty"`
	SystemSleepTimer                      int    `json:"System Sleep Timer,omitempty"`
	WakeOnLAN                             string `json:"Wake On LAN,omitempty"`
}

// ScheduledEvent represents a scheduled power event.
type ScheduledEvent struct {
	AppPID       int    `json:"appPID,omitempty"`
	EventType    string `json:"eventtype,omitempty"`
	ScheduledBy  string `json:"scheduledby,omitempty"`
	Time         string `json:"time,omitempty"`
	UserVisible  bool   `json:"UserVisible,omitempty"`
}

// ScheduledEventsInfo represents scheduled events information.
type ScheduledEventsInfo struct {
	Name   string           `json:"_name"`
	Items  []ScheduledEvent `json:"_items,omitempty"`
}

// EventsInfo represents power events information.
type EventsInfo struct {
	Name  string                `json:"_name"`
	Items []ScheduledEventsInfo `json:"_items,omitempty"`
}

// DataTypeItem represents the structure of SPPowerDataType.
type DataTypeItem struct {
	Name                           string           `json:"_name"`
	SppowerBatteryChargeInfo      BatteryChargeInfo `json:"sppower_battery_charge_info,omitempty"`
	SppowerBatteryHealthInfo      BatteryHealthInfo `json:"sppower_battery_health_info,omitempty"`
	SppowerBatteryModelInfo       BatteryModelInfo  `json:"sppower_battery_model_info,omitempty"`
	ACPower                       PowerSettings     `json:"AC Power,omitempty"`
	BatteryPower                  PowerSettings     `json:"Battery Power,omitempty"`
	SppowerUPSInstalled           string            `json:"sppower_ups_installed,omitempty"`
	SppowerBatteryChargerConnected string            `json:"sppower_battery_charger_connected,omitempty"`
	SppowerBatteryIsCharging      string            `json:"sppower_battery_is_charging,omitempty"`
	Items                         []EventsInfo      `json:"_items,omitempty"`
}

// DataType holds the parsed system profiler data for SPPowerDataType.
var DataType *common.DataType[DataTypeItem]

var (
	initOnce sync.Once
	initErr  error
)

// Initialize ensures the DataType is initialized exactly once
func Initialize() error {
	initOnce.Do(func() {
		data, err := common.NewData[DataTypeItem](common.SPPowerDataType)
		if err != nil {
			initErr = fmt.Errorf("failed to initialize power data: %w", err)
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

