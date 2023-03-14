// Package leap16 implements the LEAP16 platform.
//
// The LEAP16 platform is a 16-bit platform. It has 16-bit instructions, 16-bit
// registers, and 16-bit memory addressed using 16-bit addresses.
//
// Memory is 64 Ki16b. mFFFF and below is used as a stack. Registers are from r0
// to rF. r0 is zero, rE is the stack pointer, rF is the instruction pointer.
//
// The LEAP16 platform has instructions as follows:
//
//	LOAD ra4 or4 rd4 (0) // Load from memory ra+or, store into register rd
//	STOR ra4 or4 rs4 (E) // Store into memory ra+or, load from register rs
//	ADD  rx4 ry4 rd4 (2) // Add rx to ry, store in rd
//	SUB  rx4 ry4 rd4 (3) // Subtract ry from rx, store in rd
//	AND  rx4 ry4 rd4 (4) // And rx with ry, store in rd
//	OR   rx4 ry4 rd4 (5) // Or rx with ry, store in rd
//	SL   rx4 sa4 rd4 (6) // Shift-left rx by sa, store in rd
//	SR   rx4 sa4 rd4 (7) // Shift-right rx by sa, store in rd
//	LEAP ra4     or8 (8) // Leap to rx+or
//	LL   ra4     or8 (A) // Leap to rx+or and link
//	RL           012 (B) // Return from leap and link
//	LEQ  rx4 ry4 oi4 (C) // Leap to rF+oi if rx == ry
//	LLT  rx4 ry4 oi4 (D) // Leap to rF+oi if rx < ry
//	HALT         012 (F) // Halt execution
//
// The two unused opcodes are 1 and 9, and these non-instructions are skipped.
package leap16

// LEAP16 is an instance of the LEAP16 platform. It has 64 Ki16b of memory, 16
// registers, and a cycle counter. It can run programs loaded into memory.
type LEAP16 struct {
	// Memory
	Memory [0x10000]uint16
	// Registers
	Registers [0x10]uint16
	// Cycle counter
	C uint16
}

// NewLEAP16 returns a new instance of the LEAP16 platform.
func NewLEAP16() (l16 *LEAP16) {
	return &LEAP16{}
}

// LoadProgram loads a program into memory at address 0x0000. The program should
// not be longer than 0x10000.
func (l16 *LEAP16) LoadProgram(program []uint16) {
	copy(l16.Memory[:], program)
}

// Run runs the program loaded into memory from where the instruction pointer
// is. It will run until it hits a HALT instruction.
func (l16 *LEAP16) Run() {
	for !l16.Cycle() {
	}
}

// Cycle runs a single cycle of the program. It will return an error if it hits
// a HALT.
func (l16 *LEAP16) Cycle() (halt bool) {
	// Fetch the instruction
	instruction := l16.Memory[l16.Registers[0xF]]
	// Increment the instruction pointer
	l16.Registers[0xF]++
	// Increment the cycle counter
	l16.C++
	// Decode and execute the instruction
	opcode := instruction >> 12
	switch opcode {
	case 0x0: // LOAD
		ra := (instruction >> 8) & 0b1111
		or := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Memory[l16.Registers[ra]+or]
	case 0xE: // STOR
		ra := (instruction >> 8) & 0b1111
		or := (instruction >> 4) & 0b1111
		rs := (instruction >> 0) & 0b1111
		l16.Memory[l16.Registers[ra]+or] = l16.Registers[rs]
	case 0x2: // ADD
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] + l16.Registers[ry]
	case 0x3: // SUB
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] - l16.Registers[ry]
	case 0x4: // AND
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] & l16.Registers[ry]
	case 0x5: // OR
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] | l16.Registers[ry]
	case 0x6: // SL
		rx := (instruction >> 8) & 0b1111
		sa := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] << sa
	case 0x7: // SR
		rx := (instruction >> 8) & 0b1111
		sa := (instruction >> 4) & 0b1111
		rd := (instruction >> 0) & 0b1111
		l16.Registers[rd] = l16.Registers[rx] >> sa
	case 0x8: // LEAP
		ra := (instruction >> 8) & 0b1111
		or := (instruction >> 0) & 0b11111111
		l16.Registers[0xF] = l16.Registers[ra] + uint16(int8(or))
	case 0xA: // LL
		ra := (instruction >> 8) & 0b1111
		or := (instruction >> 0) & 0b11111111
		l16.Registers[0xE]--
		l16.Memory[l16.Registers[0xE]] = l16.Registers[0xF]
		l16.Registers[0xF] = l16.Registers[ra] + uint16(int8(or))
	case 0xB: // RL
		l16.Registers[0xF] = l16.Memory[l16.Registers[0xE]]
		l16.Registers[0xE]++
	case 0xC: // LEQ
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		oi := (instruction >> 0) & 0b1111
		if l16.Registers[rx] == l16.Registers[ry] {
			l16.Registers[0xF] += uint16(int8(oi) << 4 >> 4)
		}
	case 0xD: // LLT
		rx := (instruction >> 8) & 0b1111
		ry := (instruction >> 4) & 0b1111
		oi := (instruction >> 0) & 0b1111
		if l16.Registers[rx] < l16.Registers[ry] {
			l16.Registers[0xF] += uint16(int8(oi) << 4 >> 4)
		}
	case 0xF: // HALT
		halt = true
	default: // 0x1, 0x9
		// do nothing
	}
	return
}

// Reset resets the LEAP16 to its initial state (all zero).
func (l16 *LEAP16) Reset() {
	l16.Memory = [0x10000]uint16{}
	l16.Registers = [0x10]uint16{}
	l16.C = 0
}
