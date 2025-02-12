package common

import (
  "os/exec"
)

func ExecuteSPCommand(sp_type string) ([]byte, error) { 
  // "SPNVMeDataType"
  cmd := exec.Command("system_profiler", sp_type, "-json")
  output, err := cmd.Output()
  if err != nil {
    return nil, err
  }
  return output, err
}
