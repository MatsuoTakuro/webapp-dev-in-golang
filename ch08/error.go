package ch08

import (
	"errors"
	"fmt"
)

type ErrCode int
type MyError struct {
	Code ErrCode
	Err  error
}

func (e *MyError) Error() string {
	return fmt.Sprintf("code: %d", e.Code)
}

var _ error = (*MyError)(nil)

func (e *MyError) Unwrap() error {
	return errors.Unwrap(e.Err)
}

func (e *MyError) As(target interface{}) bool {
	return errors.As(e.Err, target)
}

func (e *MyError) Is(target error) bool {
	return errors.Is(e.Err, target)
}
