package register

import (
	"fmt"

	"github.com/akyoto/color"
	"github.com/akyoto/q/build/log"
)

// Register represents a single CPU register.
type Register struct {
	ID       ID
	Name     string
	usedBy   fmt.Stringer
	hasValue bool
}

// Use marks the register as used by the given object.
func (register *Register) Use(obj fmt.Stringer) error {
	if obj == nil {
		panic("register.Use parameter cannot be nil")
	}

	if register.usedBy != nil {
		return &ErrAlreadyInUse{register, register.usedBy}
	}

	register.usedBy = obj
	return nil
}

// ForceUse marks the register as used by the given object and cannot fail.
func (register *Register) ForceUse(obj fmt.Stringer) {
	if obj == nil {
		panic("register.ForceUse parameter cannot be nil")
	}

	register.usedBy = obj
}

// User returns the user of the register.
func (register *Register) User() fmt.Stringer {
	return register.usedBy
}

// UserString returns the user as a string.
// If the user is nil, it returns "?".
func (register *Register) UserString() string {
	if register.usedBy != nil {
		return register.usedBy.String()
	}

	return "?"
}

// Free frees the register so that it can be used for new calculations.
func (register *Register) Free() {
	register.usedBy = nil
	register.hasValue = false
}

// IsFree returns true if the register is not in use.
func (register *Register) IsFree() bool {
	return register.usedBy == nil
}

// Assign marks the register as assigned which means it holds a value.
func (register *Register) Assign() {
	register.hasValue = true
}

// IsEmpty returns true if the register has no value.
func (register *Register) IsEmpty() bool {
	return !register.hasValue
}

// String returns a human-readable representation of the register.
func (register *Register) String() string {
	return register.StringWithUser(register.UserString())
}

// StringWithUser returns a human-readable representation of the register.
func (register *Register) StringWithUser(usedBy string) string {
	return fmt.Sprintf("%s%s%v", register.Name, log.FaintColor.Sprint("="), color.GreenString(usedBy))
}
