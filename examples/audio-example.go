package main

import (
	"fmt"
	"log"

	"github.com/samburba/go-system-profiler/v2/type/audio"
)

func main() {
	// Get the audio data (initializes if needed)
	data, err := audio.GetDataType()
	if err != nil {
		log.Fatalf("Failed to get audio data: %v", err)
	}

	// Print the data in a formatted JSON string
	fmt.Printf("%s\n", data.String())

	// Iterate over the audio devices and print details
	for _, device := range data.Item {
		fmt.Printf("Device Name: %s\n", device.Name)
		fmt.Printf("Input Source: %s\n", device.CoreaudioInputSource)
		fmt.Printf("Manufacturer: %s\n", device.CoreaudioDeviceManufacturer)
		fmt.Printf("Sample Rate: %d\n", device.CoreaudioDeviceSrate)
		fmt.Printf("Transport: %s\n", device.CoreaudioDeviceTransport)
	}
}
