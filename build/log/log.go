package log

import (
	"log"
	"os"

	"github.com/akyoto/color"
)

var (
	// Info is used for general info messages.
	Info = log.New(os.Stdout, "", 0)

	// Error is used for error messages.
	Error = log.New(os.Stderr, "", 0)
)

var (
	// FaintColor is the color used for printing faint messages.
	FaintColor = color.New(color.Faint)

	// CommentColor is the color used for printing comments.
	CommentColor = color.New(color.FgYellow)
)
