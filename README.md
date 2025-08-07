# go-system-profiler - Go System Profiler

## Overview
go-system-profiler is a Go library that retrieves and structures system information using macOS's `system_profiler` command. It provides a structured and JSON-compatible output to facilitate system data parsing and analysis.

## Features
- Retrieves system information using `system_profiler`
- Supports multiple data types (e.g., Audio, Hardware, Software, Network, etc.)
- Provides JSON-formatted structured output
- Uses generics for flexible data handling
- Modular design for easy expansion
- Thread-safe initialization with `sync.Once`
- Comprehensive test coverage for implemented types

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
| Type                 | Implementation Status | Test Coverage |
|----------------------|----------------------|---------------|
| `audio/`           | ✅ Implemented | ✅ Tested |
| `hardware/`        | ✅ Implemented | ✅ Tested |
| `software/`        | ✅ Implemented | ✅ Tested |
| `network/`         | ✅ Implemented | ✅ Tested |
| `usb/`             | ✅ Implemented | ✅ Tested |
| `bluetooth/`       | ✅ Implemented | ✅ Tested |
| `applications/`    | ✅ Implemented | ⚠️ Basic |
| `memory/`          | ✅ Implemented | ⚠️ Basic |
| `power/`           | ✅ Implemented | ⚠️ Basic |
| `storage/`         | ✅ Implemented | ⚠️ Basic |
| `nvme/`            | ✅ Implemented | ⚠️ Basic |
| `displays/`        | ✅ Implemented | ⚠️ Basic |
| *Other types*      | ✅ Basic Structure | ❌ No Tests |

## Data Structure Notes
The library currently supports two types of system_profiler data structures:

1. **Items-based structure** (like Audio): Data with `_items` array containing multiple devices/items
2. **Flat structure** (like Hardware, Software): Single object with direct properties

Types with items-based structure (like Audio) provide full access to their fields, while types with flat structure currently only expose the basic `Name` field due to the common package design.

## Usage

### Initialization Pattern
All type packages use a thread-safe initialization pattern with `sync.Once`. You have two ways to access the data:

#### Method 1: Using `Initialize()` and `DataType` directly
```go
package main

import (
	"fmt"
	"log"
	"github.com/samburba/go-system-profiler/v2/type/audio"
)

func main() {
	// Initialize the data (thread-safe, only happens once)
	if err := audio.Initialize(); err != nil {
		log.Fatalf("Failed to initialize audio data: %v", err)
	}

	// Access the data directly
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

#### Method 2: Using `GetDataType()` (recommended)
```go
package main

import (
	"fmt"
	"log"
	"github.com/samburba/go-system-profiler/v2/type/audio"
)

func main() {
	// Get the data type (initializes if needed)
	data, err := audio.GetDataType()
	if err != nil {
		log.Fatalf("Failed to get audio data: %v", err)
	}

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

### Working with Different Data Types

#### Audio Data (Items-based structure)
```go
package main

import (
	"fmt"
	"log"
	"github.com/samburba/go-system-profiler/v2/type/audio"
)

func main() {
	data, err := audio.GetDataType()
	if err != nil {
		log.Fatalf("Failed to get audio data: %v", err)
	}

	// Access individual audio devices
	for _, device := range data.Item {
		fmt.Printf("Audio Device: %s\n", device.Name)
		fmt.Printf("  Input Source: %s\n", device.CoreaudioInputSource)
		fmt.Printf("  Output Source: %s\n", device.CoreaudioOutputSource)
		fmt.Printf("  Manufacturer: %s\n", device.CoreaudioDeviceManufacturer)
		fmt.Printf("  Sample Rate: %d\n", device.CoreaudioDeviceSrate)
	}
}
```

#### Hardware Data (Flat structure)
```go
package main

import (
	"fmt"
	"log"
	"github.com/samburba/go-system-profiler/v2/type/hardware"
)

func main() {
	data, err := hardware.GetDataType()
	if err != nil {
		log.Fatalf("Failed to get hardware data: %v", err)
	}

	// Access hardware information
	for _, item := range data.Item {
		fmt.Printf("Hardware Component: %s\n", item.Name)
		// Note: Additional fields need to be added to the DataTypeItem struct
	}
}
```

#### Network Data
```go
package main

import (
	"fmt"
	"log"
	"github.com/samburba/go-system-profiler/v2/type/network"
)

func main() {
	data, err := network.GetDataType()
	if err != nil {
		log.Fatalf("Failed to get network data: %v", err)
	}

	// Access network interfaces
	for _, interface := range data.Item {
		fmt.Printf("Network Interface: %s\n", interface.Name)
		// Access interface-specific fields
	}
}
```

## Examples
See [audio-example.go](examples/audio-example.go) for a complete example usage of the `audio` package.

## Testing
Run tests for all implemented types:
```bash
go test ./type/... -v
```

Run tests for a specific type:
```bash
go test ./type/audio/... -v
```

## Building
Build the entire project:
```bash
go build ./...
```

## Requirements
- macOS (uses `system_profiler` command)
- Go 1.18+ (for generics support)

## Thread Safety
All type packages use `sync.Once` to ensure thread-safe initialization. The `Initialize()` function can be called multiple times safely, and the actual initialization only happens once. The `GetDataType()` function is the recommended way to access data as it handles initialization automatically.

## CI/CD and Releases

This project uses GitHub Actions for automated testing and releases:

### Automated Releases
- **Main Branch:** Pushes to `main` automatically create versioned releases (e.g., `v1.2.3`)
- **Nightly Builds:** Daily automated releases with `-nightly.YYYYMMDD` suffix
- **Pull Requests:** Comprehensive testing and code quality checks

### Test Matrix
- **macOS Latest:** Go 1.21, 1.22, 1.23
- **macOS 13:** Go 1.22, 1.23  
- **macOS 12:** Go 1.22, 1.23

### Quality Checks
- Unit tests for all type packages
- Code formatting with `gofmt`
- Static analysis with `go vet`
- Race condition detection
- Build verification

For more details, see [`.github/workflows/README.md`](.github/workflows/README.md).


