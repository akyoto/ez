package build

import (
	"strconv"

	"github.com/akyoto/q/build/errors"
	"github.com/akyoto/q/build/token"
)

// AssignArrayElement assigns a value to an array element.
func (state *State) AssignArrayElement(tokens []token.Token, operatorPos token.Position) error {
	left := tokens[:operatorPos]
	right := tokens[operatorPos+1:]
	suffix := left[1:]

	if suffix[0].Kind != token.ArrayStart || suffix[len(suffix)-1].Kind != token.ArrayEnd {
		return errors.New(errors.MissingArrayIndex)
	}

	if left[0].Kind != token.Identifier {
		return errors.New(errors.ExpectedVariable)
	}

	arrayName := left[0].Text()
	array := state.scopes.Get(arrayName)

	valueString := right[0].Text()
	value, err := strconv.Atoi(valueString)

	if err != nil {
		return err
	}

	indexTokens := suffix[1 : len(suffix)-1]

	if len(indexTokens) > 1 || indexTokens[0].Kind != token.Number {
		return errors.New(errors.NotImplemented)
	}

	indexString := indexTokens[0].Text()
	index, err := strconv.Atoi(indexString)

	if err != nil {
		return err
	}

	state.assembler.StoreNumber(array.Register(), byte(index), byte(1), uint64(value))
	return nil
}
