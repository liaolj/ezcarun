package main

import (
	"fmt"

	"github.com/liaolj/ebridge"
)

func main() {
	// d := ebridge.Edemo()
	value, err := ebridge.LongGet("epicsHost:ai1")
	if err != nil {
		fmt.Printf("result is %d\n", value)
	}
}
