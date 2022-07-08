package build

import (
	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// AssignArrayElement assigns a value to an array element.
func (state *State) AssignArrayElement(tokens []token.Token, operatorPos token.Position) error {
	left := tokens[:operatorPos]
	suffix := left[1:]

	if suffix[0].Kind == token.ArrayStart && suffix[len(suffix)-1].Kind == token.ArrayEnd {
		indexTokens := suffix[1 : len(suffix)-1]

		if indexTokens[0].Kind != token.Number {
			return errors.New(errors.NotImplemented)
		}

		return errors.New(errors.NotImplemented)
	}

	return nil
}
