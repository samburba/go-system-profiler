package nvme

import (
	"github.com/samburba/go-system-profiler/v2/internal/common"
)

const ThisDataType = "SPNVMeDataType"

type DataTypeItem struct {
	Name              string `json:"_name"`
	BsdName           string `json:"bsd_name"`
	DetachableDrive   string `json:"detachable_drive"`
	DeviceModel       string `json:"device_model"`
	DeviceRevision    string `json:"device_revision"`
	DeviceSerial      string `json:"device_serial"`
	PartitionMapType  string `json:"partition_map_type"`
	RemovableMedia    string `json:"removable_media"`
	Size              string `json:"size"`
	SizeInBytes       int64  `json:"size_in_bytes"`
	SmartStatus       string `json:"smart_status"`
	SpnvmeTrimSupport string `json:"spnvme_trim_support"`
	Volumes           []struct {
		Name        string `json:"_name"`
		BsdName     string `json:"bsd_name"`
		Iocontent   string `json:"iocontent"`
		Size        string `json:"size"`
		SizeInBytes int    `json:"size_in_bytes"`
	} `json:"volumes"`
}

var NVMeData *common.Data = common.NewData(ThisDataType)
