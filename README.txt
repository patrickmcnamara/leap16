================================== LEAP16 ===================================

The LEAP16 platform is a 16-bit platform. It has 16-bit instructions, 16-bit
registers, and 16-bit memory addressed using 16-bit addresses.

Memory is 64 Ki16b. mFFFF and below is used as a stack. Registers are from r0
to rF. r0 is zero, rE is the stack pointer, rF is the instruction pointer.

The LEAP16 platform has instructions as follows:

LOAD ra4 or4 rd4 (0) // Load from memory ra+or, store into register rd
STOR ra4 or4 rs4 (E) // Store into memory ra+or, load from register rs
ADD  rx4 ry4 rd4 (2) // Add rx to ry, store in rd
SUB  rx4 ry4 rd4 (3) // Subtract ry from rx, store in rd
AND  rx4 ry4 rd4 (4) // And rx with ry, store in rd
OR   rx4 ry4 rd4 (5) // Or rx with ry, store in rd
SL   rx4 sa4 rd4 (6) // Shift-left rx by sa, store in rd
SR   rx4 sa4 rd4 (7) // Shift-right rx by sa, store in rd
LEAP ra4     or8 (8) // Leap to rx+or
LL   ra4     or8 (A) // Leap to rx+or and link
RL           012 (B) // Return from leap and link
LEQ  rx4 ry4 oi4 (C) // Leap to rF+oi if rx == ry
LLT  rx4 ry4 oi4 (D) // Leap to rF+oi if rx < ry
HALT         012 (F) // Halt execution

The two unused opcodes are 1 and 9, and these non-instructions are skipped.

=============================================================================

Included in this repo is a reference implementation of the LEAP16 in Go.

=============================================================================

Copyright (c) 2023 by Patrick McNamara. Released under the MIT License.
