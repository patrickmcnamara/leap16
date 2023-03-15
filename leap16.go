// Package leap16 implements the LEAP16 computer architecture.
//
// LEAP16 is a 16-bit computer architecture. It has 16-bit instructions, 16-bit
// registers, and 16-bit I/O addressed using 16-bit addresses. In this version,
// all I/O is memory.
//
// Memory is from m0000 to mFFFF. m0000 and above is used as a stack. Registers
// are from r0 to rF. rE is the stack pointer, rF is the instruction pointer. r0
// is 0 by convention.
//
// The LEAP16 instruction set is as follows:
//
//	R     ra  or  rd    // Read from ra+or, store to rd
//	W     ra  or  rs    // Write to ra+or, load from rs
//	ADD   rx  ry  rd    // Add rx to ry, store to rd
//	SUB   rx  ry  rd    // Subtract ry from rx, store to rd
//	AND   rx  ry  rd    // And rx with ry, store to rd
//	OR    rx  ry  rd    // Or rx with ry, store to rd
//	SL    rx  sa  rd    // Shift-left rx by sa, store to rd
//	SR    rx  sa  rd    // Shift-right rx by sa, store to rd
//	LEAP  ra      or    // Leap to rx+or
//	LL    ra      or    // Leap to rx+or and link
//	RL                  // Return from leap and link
//	LEQ   rx  ry  oi    // Leap to rF+oi if rx == ry
//	LLT   rx  ry  oi    // Leap to rF+oi if rx < ry
//	HALT                // Halt execution
//
// The opcodes for instructions are the index of that instruction, with
// exceptions that W is E, and 1 and 9 are skipped. R is 0 and HALT is F, etc.
package leap16

// LEAP16 is an instance of a LEAP16 computer. It has 16 registers,  64 Ki of
// memory, and a cycle counter. It can run programs loaded into memory.
type LEAP16 struct {
	Registers [0x10]uint16
	Memory    [0x10000]uint16
	C         uint64
}

// NewLEAP16 returns a new instance of a LEAP16 computer.
func NewLEAP16() (l16 *LEAP16) {
	return &LEAP16{}
}

// LoadProgram loads a program into the LEAP16 computer memory at address 0000.
// The program should not be longer than FFFF.
func (l16 *LEAP16) LoadProgram(program []uint16) {
	copy(l16.Memory[:], program)
}

// Run runs the LEAP16 computer until it halts.
func (l16 *LEAP16) Run() {
	for !l16.Cycle() {
	}
}

// Cycle runs a single cycle of the LEAP16 computer. This:
//  1. Fetches the instruction from memory
//  2. Increments the instruction pointer and the cycle counter
//  3. Decodes the instruction
//  4. Executes the instruction
//  5. Returns true if the executed instruction was HALT, false otherwise
func (l16 *LEAP16) Cycle() (halt bool) {
	// Fetch the instruction
	instruction := l16.Memory[l16.Registers[0xF]]
	// Increment the instruction pointer
	l16.Registers[0xF]++
	// Increment the cycle counter
	l16.C++
	// Decode the instruction
	o0 := (instruction >> 0xC) & 0xF  // 4-bit opcode
	f3 := (instruction >> 0x8) & 0xF  // 4-bit operands
	f7 := (instruction >> 0x4) & 0xF  //
	fB := (instruction >> 0x0) & 0xF  //
	e7 := (instruction >> 0x0) & 0xFF // 8-bit operand
	// Execute the instruction
	switch o0 {
	case OPCODE_R:
		l16.Registers[fB] = l16.Memory[l16.Registers[f3]+f7]
	case OPCODE_W:
		l16.Memory[l16.Registers[f3]+f7] = l16.Registers[fB]
	case OPCODE_ADD:
		l16.Registers[fB] = l16.Registers[f3] + l16.Registers[f7]
	case OPCODE_SUB:
		l16.Registers[fB] = l16.Registers[f3] - l16.Registers[f7]
	case OPCODE_AND:
		l16.Registers[fB] = l16.Registers[f3] & l16.Registers[f7]
	case OPCODE_OR:
		l16.Registers[fB] = l16.Registers[f3] | l16.Registers[f7]
	case OPCODE_SL:
		l16.Registers[fB] = l16.Registers[f3] << f7
	case OPCODE_SR:
		l16.Registers[fB] = l16.Registers[f3] >> f7
	case OPCODE_LEAP:
		l16.Registers[0xF] = l16.Registers[f3] + uint16(int8(e7))
	case OPCODE_LL:
		l16.Memory[l16.Registers[0xE]] = l16.Registers[0xF]
		l16.Registers[0xE]++
		l16.Registers[0xF] = l16.Registers[f3] + uint16(int8(e7))
	case OPCODE_RL:
		l16.Registers[0xE]--
		l16.Registers[0xF] = l16.Memory[l16.Registers[0xE]]
	case OPCODE_LEQ:
		if l16.Registers[f3] == l16.Registers[f7] {
			l16.Registers[0xF] += uint16(int8(fB) << 4 >> 4)
		}
	case OPCODE_LLT:
		if l16.Registers[f3] < l16.Registers[f7] {
			l16.Registers[0xF] += uint16(int8(fB) << 4 >> 4)
		}
	case OPCODE_HALT:
		halt = true
	default: // 1 or 9 (undefined opcode)
		// do nothing
	}
	return
}

// Reset resets the LEAP16 computer to the initial state (all zero).
func (l16 *LEAP16) Reset() {
	l16.Memory = [0x10000]uint16{}
	l16.Registers = [0x10]uint16{}
	l16.C = 0
}
