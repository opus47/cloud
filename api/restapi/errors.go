package restapi

import (
	"fmt"
)

type errort interface {
	error
	Type() ErrorType
}

type ErrorType int

const (
	User ErrorType = iota
	System
)

type Error struct {
	Typ     ErrorType
	Message string
}

func (x Error) Error() string {
	return x.Message
}

func (x Error) Type() ErrorType {
	return x.Typ
}

func SysError(s string, args ...interface{}) Error {
	return Error{
		System,
		fmt.Sprintf(s, args...),
	}
}

func UserError(s string, args ...interface{}) Error {
	return Error{
		User,
		fmt.Sprintf(s, args...),
	}
}
