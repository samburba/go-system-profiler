package nvme

import (
	"encoding/json"
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

type DataType struct {
  Item []DataTypeItem `json:"_items"`
  Name string `json:"_name"`
}

type NVMeData struct{
  SPNVMeDataType []DataType `json:"SPNVMeDataType"`
} 

var nvmeData NVMeData 

func New() *NVMeData {
  if len(nvmeData.SPNVMeDataType) == 0 {
    nvmeData = executeSPCommand()
  }

    return &nvmeData
}

func executeSPCommand() NVMeData {
	output, err := common.ExecuteSPCommand(ThisDataType)
	if err != nil {
		panic(err)
	}
  var data NVMeData
	err = json.Unmarshal(output, &data)
	if err != nil {
		panic(err)
	}
	return data
}
