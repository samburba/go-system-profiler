package common

import (
  "encoding/json"
  "os/exec"
)


type Data struct {
  DataType []struct {
    Item []any  `json:"_items"`
    Name string `json:"_name"`
  } `json:"SPNVMeDataType"`
}


func NewData(sp_type string) *Data {
  d, err := executeSPCommand(sp_type)
  if err != nil {
    panic(err) //don't want to pass further down
  }
  return d
}


func executeSPCommand(sp_type string) (*Data, error) { 
  cmd := exec.Command("system_profiler", sp_type, "-json")
  output, err := cmd.Output()
  if err != nil {
    return nil, err
  }
  var data Data
	err = json.Unmarshal(output, &data)
	if err != nil {
    return nil, err
	}
	return &data, nil
}
