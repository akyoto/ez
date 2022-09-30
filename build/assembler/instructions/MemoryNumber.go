package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// MemoryNumber is used for instructions requiring a memory and a number operand.
type MemoryNumber struct {
	Base
	Destination *register.Register
	Number      uint64
	UsedBy      string
	Offset      byte
	ByteCount   byte
}

// Exec writes the instruction to the final assembler.
func (instr *MemoryNumber) Exec(a *asm.Assembler) {
	start := a.Position()

	switch instr.Mnemonic {
	case mnemonics.STORE:
		a.StoreNumber(instr.Destination.Name, instr.Offset, instr.ByteCount, instr.Number)

	default:
		panic("This should never happen!")
	}

	instr.size = byte(a.Position() - start)
}

// String implements the string serialization.
func (instr *MemoryNumber) String() string {
	return fmt.Sprintf("%s %dB [%v+%d], %d", mnemonicColor.Sprint(instr.Mnemonic), instr.ByteCount, instr.Destination.StringWithUser(instr.UsedBy), instr.Offset, instr.Number)
}
