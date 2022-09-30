package instructions

import (
	"github.com/akyoto/asm"
	"github.com/akyoto/q/build/log"
)

// AddComment is used for instructions that add a comment.
type AddComment struct {
	Comment string
}

// Exec writes the instruction to the final assembler.
func (instr *AddComment) Exec(a *asm.Assembler) {
	// Not applicable.
}

// Name returns the empty string.
func (instr *AddComment) Name() string {
	return "COMMENT"
}

// SetName sets the mnemonic.
func (instr *AddComment) SetName(mnemonic string) {
	// Not applicable.
}

// Size returns the number of bytes consumed for the instruction.
func (instr *AddComment) Size() byte {
	return 0
}

// String implements the string serialization.
func (instr *AddComment) String() string {
	return log.CommentColor.Sprint(instr.Comment)
}
