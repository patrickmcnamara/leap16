// Package l16util provides utility functions for the leap16 package.
package l16util

import (
	"fmt"

	"github.com/patrickmcnamara/leap16"
)

// Dump dumps the LEAP16 computer registers, memory, and cycle count to stdout.
func Dump(l16 *leap16.LEAP16, me uint16) {
	// Registers
	for i := 0; i < 0x10; i++ {
		fmt.Printf("r%01X:     %04X\n", i, l16.Registers[i])
	}
	// I/O
	for i := 0; i < int(me); i++ {
		fmt.Printf("io%04X:  %04X\n", i, l16.Memory[i])
	}
	// Cycle count
	fmt.Printf("c:      %08X\n", l16.C)
}
