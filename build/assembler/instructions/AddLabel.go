package instructions

import (
	"fmt"

	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/log"
)

// AddLabel is used for instructions that add a label.
type AddLabel struct {
	Label string
}

// Exec writes the instruction to the final assembler.
func (instr *AddLabel) Exec(a *asm.Assembler) {
	a.AddLabel(instr.Label)
}

// Name returns the empty string.
func (instr *AddLabel) Name() string {
	return "LABEL"
}

// SetName sets the mnemonic.
func (instr *AddLabel) SetName(mnemonic string) {
	// Not applicable.
}

// Size returns the number of bytes consumed for the instruction.
func (instr *AddLabel) Size() byte {
	return 0
}

// String implements the string serialization.
func (instr *AddLabel) String() string {
	return fmt.Sprintf("%s:", log.FaintColor.Sprint(instr.Label))
}
