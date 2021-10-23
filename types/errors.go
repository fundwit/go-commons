package types

import (
	"fmt"
)

type ErrInvalidParameter struct {
	Parameter interface{}
	Cause     error
}

func (e *ErrInvalidParameter) Error() string {
	return fmt.Sprintf("invalid parameter: %#v", e.Parameter)
}

func (e *ErrInvalidParameter) Unwrap() error {
	return e.Cause
}
