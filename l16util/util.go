// Package l16util provides utility functions for the leap16 package.
package l16util

import (
	"fmt"

	"github.com/patrickmcnamara/leap16"
)

// Dump dumps the contents of the LEAP16 registers, memory, and stack (optional)
// to os.Stdout.
func Dump(l16 *leap16.LEAP16, me uint16, ds bool) {
	// Registers
	for i := 0; i < 16; i++ {
		fmt.Printf("r%01X:     %04X\n", i, l16.Registers[i])
	}
	// Memory
	for i := 0; i < int(me); i++ {
		fmt.Printf("m%04X:  %04X\n", i, l16.Memory[i])
	}
	// Stack (optionally)
	if ds && l16.Registers[0xE] != 0 {
		for i := int(l16.Registers[0xE]); i < 0x10000; i++ {
			fmt.Printf("sm%04X: %04X\n", i, l16.Memory[i])
		}
	}
	// Cycle count
	fmt.Printf("c:      %04X\n", l16.C)
}
