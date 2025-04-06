==================================== LEAP16 ====================================

LEAP16 is a 16-bit computer architecture. It has 16-bit instructions, 16-bit
registers, and 16-bit I/O addressed using 16-bit addresses.

I/O is from io0000 to ioFFFF. io0000 and above is used as a stack. Registers are
from r0 to rF. rE is the stack pointer, rF is the instruction pointer. r0 is 0
by convention.

The LEAP16 instruction set is as follows:

	R     ra  or  rd    // Read from ra+or, store to rd
	W     ra  or  rs    // Write to ra+or, load from rs
	ADD   rx  ry  rd    // Add rx to ry, store to rd
	SUB   rx  ry  rd    // Subtract ry from rx, store to rd
	AND   rx  ry  rd    // And rx with ry, store to rd
	OR    rx  ry  rd    // Or rx with ry, store to rd
	SL    rx  sa  rd    // Shift-left rx by sa, store to rd
	SR    rx  sa  rd    // Shift-right rx by sa, store to rd
	LEAP  ra      or    // Leap to ra+or
	LL    ra      or    // Leap to ra+or and link
	RLL                 // Return from leap and link
	LEQ   rx  ry  oi    // Leap to rF+oi if rx = ry
	LLT   rx  ry  oi    // Leap to rF+oi if rx < ry
	HALT                // Halt execution

	ra = register address (register index)
	rd = register destination
	rs = register source
	rx = register X
	ry = register Y
	or = offset from register (4b/8b immediate)
	oi = offset from instruction pointer (4b immediate)
	sa = shift amount (4b immediate)

The opcodes for instructions are the index of that instruction, with exceptions
that W is E, and 1 and 9 are skipped. R is 0 and HALT is F, etc.

================================================================================

Included in this repo is a reference implementation of LEAP16 in Go.

================================================================================

Copyright (c) 2023 by Patrick McNamara. Released under the MIT License.
