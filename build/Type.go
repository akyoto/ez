package build

import (
	"bytes"

	"github.com/akyoto/q/build/token"
)

// TypeNameFromTokens converts tokens to the type name they represent.
func TypeNameFromTokens(tokens []token.Token) string {
	typeName := bytes.Buffer{}

	for _, token := range tokens {
		typeName.Write(token.Bytes)
	}

	return typeName.String()
}
