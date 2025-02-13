package main

import (
	"fmt"
	"github.com/samburba/go-system-profiler/v2/audio"
)

func main() {
	a := audio.Data
	fmt.Printf("%s\n", a.String())
}
