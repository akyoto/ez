package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/assembler/mnemonics"
	"github.com/akyoto/q/build/register"
)

// RegisterRegister is used for instructions requiring 2 register operands.
type RegisterRegister struct {
	Base
	Destination *register.Register
	Source      *register.Register
	UsedBy1     string
	UsedBy2     string
}

// Exec writes the instruction to the final assembler.
func (instr *RegisterRegister) Exec(a *asm.Assembler) {
	start := a.Position()

	switch instr.Mnemonic {
	case mnemonics.MOV:
		a.MoveRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case mnemonics.CMP:
		a.CompareRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case mnemonics.ADD:
		a.AddRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case mnemonics.SUB:
		a.SubRegisterRegister(instr.Destination.Name, instr.Source.Name)

	case mnemonics.MUL:
		a.MulRegisterRegister(instr.Destination.Name, instr.Source.Name)
	}

	instr.size = byte(a.Position() - start)
}

// String implements the string serialization.
func (instr *RegisterRegister) String() string {
	return fmt.Sprintf("%s %v, %v", mnemonicColor.Sprint(instr.Mnemonic), instr.Destination.StringWithUser(instr.UsedBy1), instr.Source.StringWithUser(instr.UsedBy2))
}
