package instruction

import (
	"strings"

	"github.com/akyoto/q/build/token"
)

// Instruction encapsulates a single instruction inside a function.
// Instructions can be variable assignments, function calls or keywords.
type Instruction struct {
	Kind     Kind
	Tokens   []token.Token
	Position token.Position
}

// String implements the string serialization.
func (instr *Instruction) String() string {
	builder := strings.Builder{}

	for index, t := range instr.Tokens {
		if index == len(instr.Tokens)-1 {
			builder.WriteString(t.Text())
			break
		}

		if t.Kind == token.Keyword || t.Kind == token.Separator {
			builder.WriteString(t.Text())
			builder.WriteByte(' ')
			continue
		}

		if t.Kind == token.Operator && t.Bytes[0] != '.' {
			builder.WriteByte(' ')
			builder.WriteString(t.Text())
			builder.WriteByte(' ')
			continue
		}

		builder.WriteString(t.Text())
	}

	return builder.String()
}
