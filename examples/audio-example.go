package main

import (
	"fmt"
	"github.com/samburba/go-system-profiler/v2/type/audio" // Make sure this path is correct
)

func main() {
	// Create a DataType instance (make sure this path and initialization are correct)
	data := audio.DataType
	
	// Print the data in a formatted JSON string
	fmt.Printf("%s\n", data.String())

	// Iterate over the audio devices and print details
	for _, device := range data.Items() {
		fmt.Printf("Device Name: %s\n", device.Name)
		fmt.Printf("Input Source: %s\n", device.CoreaudioInputSource)
		fmt.Printf("Manufacturer: %s\n", device.CoreaudioDeviceManufacturer)
		fmt.Printf("Sample Rate: %d\n", device.CoreaudioDeviceSrate)
		fmt.Printf("Transport: %s\n", device.CoreaudioDeviceTransport)
	}
}

