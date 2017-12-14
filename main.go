package main

import (
	"fmt"

	"github.com/liaolj/ebridge"
)

func main() {
	// d := ebridge.Edemo()
	value1, err1 := ebridge.LongGet("ll")
	if err1 != nil {
		fmt.Printf("[longin]-ll is %v\n", value1)
	}

	value2, err2 := ebridge.StringGet("str")
	if err2 != nil {
		fmt.Printf("[stringin]-str is %v\n", value2)
	}
	value3, err3 := ebridge.DoubleGet("dd")
	if err3 != nil {
		fmt.Printf("[ai]-dd is %v\n", value3)
	}
	value4, err4 := ebridge.BoolGet("bb")
	if err4 != nil {
		fmt.Printf("[bi]-bb is %v\n", value4)
	}
}
