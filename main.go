package main

import (
	"fmt"

	"github.com/liaolj/ebridge"
)

func main() {
	// d := ebridge.Edemo()
	value, err := ebridge.LongGet("ll")
	if err != nil {
		fmt.Printf("[longin]-ll is %v\n", value)
	}

	value, err := ebridge.StringGet("str")
	if err != nil {
		fmt.Printf("[stringin]-str is %v\n", value)
	}
	value, err := ebridge.DoubleGet("dd")
	if err != nil {
		fmt.Printf("[ai]-dd is %v\n", value)
	}
	value, err := ebridge.BoolGet("bb")
	if err != nil {
		fmt.Printf("[bi]-bb is %v\n", value)
	}
}
