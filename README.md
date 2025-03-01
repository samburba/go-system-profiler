# go-system-profiler - Go System Profiler

## Overview
go-system-profiler is a Go library that retrieves and structures system information using macOS's `system_profiler` command. It provides a structured and JSON-compatible output to facilitate system data parsing and analysis.

## Features
- Retrieves system information using `system_profiler`
- Supports multiple data types (e.g., Audio, NVMe, USB, Network, etc.)
- Provides JSON-formatted structured output
- Uses generics for flexible data handling
- Modular design for easy expansion

## File Structure
```
/go-system-profiler/
├── README.md               # Project documentation
├── go.mod                  # Go module file
├── internal/               # Internal package for shared utilities
│   ├── common/             # Shared utilities and helpers
│   │   ├── common.go       # General helper functions
│   │   ├── spdatatype.go   # System profiler data types
├── type/                   # Modules for specific system data types
│   ├── airport/            # Airport system information module
│   ├── applications/       # Applications information module
│   ├── audio/              # Audio system information module
│   ├── bluetooth/          # Bluetooth system information module
│   ├── camera/             # Camera system information module
│   ├── cardreader/         # Card reader system information module
│   ├── configurationprofile/ # Configuration profile module
│   ├── developertools/     # Developer tools information module
│   ├── diagnostics/        # Diagnostics system information module
│   ├── disabledsoftware/   # Disabled software information module
│   ├── discburning/        # Disc burning system information module
│   ├── displays/           # Displays system information module
│   ├── ethernet/           # Ethernet system information module
│   ├── extensions/         # Extensions system information module
│   ├── fibrechannel/       # Fibre Channel system information module
│   ├── firewall/           # Firewall system information module
│   ├── firewire/           # FireWire system information module
│   ├── fonts/              # Fonts system information module
│   ├── frameworks/         # Frameworks system information module
│   ├── hardware/           # Hardware system information module
│   ├── installhistory/     # Install history information module
│   ├── international/      # International settings information module
│   ├── legacysoftware/     # Legacy software system information module
│   ├── logs/               # System logs information module
│   ├── managedclient/      # Managed client system information module
│   ├── memory/             # Memory system information module
│   ├── network/            # Network system information module
│   ├── networklocation/    # Network location information module
│   ├── networkvolume/      # Network volume information module
│   ├── nvme/               # NVMe storage information module
│   ├── parallelata/        # Parallel ATA system information module
│   ├── parallelscsi/       # Parallel SCSI system information module
│   ├── pci/                # PCI system information module
│   ├── power/              # Power management module
│   ├── prefpane/           # Preference pane system information module
│   ├── printers/           # Printers system information module
│   ├── printerssoftware/   # Printer software information module
│   ├── rawcamera/          # Raw camera system information module
│   ├── sas/                # SAS system information module
│   ├── secureelement/      # Secure element system information module
│   ├── serialata/          # Serial ATA system information module
│   ├── smartcards/         # Smart card system information module
│   ├── software/           # Software system information module
│   ├── spi/                # SPI system information module
│   ├── startupitem/        # Startup item system information module
│   ├── storage/            # Storage system information module
│   ├── syncservices/       # Sync services system information module
│   ├── thunderbolt/        # Thunderbolt system information module
│   ├── universalaccess/    # Universal access system information module
│   ├── usb/                # USB information module
```

## Implementation Status
Each row corresponds to a type subdirectory.
| Type                 | Implementation Status |
|----------------------|----------------------|
| `audio/`           | ✅ Implemented |
| `nvme/`           | ✅ Implemented |
| *Other types*      | ❌ Not Implemented |

## Examples
See [audio-example.go](examples/audio-example.go) for an example usage of the `audio` package.

### Example Usage
```go
package main

import (
	"fmt"
	"github.com/samburba/go-system-profiler/v2/type/audio"
)

func main() {
	data := audio.DataType
	fmt.Printf("%s\n", data.String())

	for _, device := range data.Item {
		fmt.Printf("Device Name: %s\n", device.Name)
		fmt.Printf("Input Source: %s\n", device.CoreaudioInputSource)
		fmt.Printf("Manufacturer: %s\n", device.CoreaudioDeviceManufacturer)
		fmt.Printf("Sample Rate: %d\n", device.CoreaudioDeviceSrate)
		fmt.Printf("Transport: %s\n", device.CoreaudioDeviceTransport)
	}
}
```


