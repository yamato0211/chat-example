package codes

import (
	"fmt"
)

type Code int

const (
	CodeUnknown Code = iota
	CodeOK
	CodeInvalidArgument
	CodeNotFound
	CodeInternal
)

func (c Code) String() string {
	switch c {
	case CodeUnknown:
		return "Unknown"
	case CodeOK:
		return "OK"
	case CodeInvalidArgument:
		return "Invalid argument"
	case CodeNotFound:
		return "Not found"
	case CodeInternal:
		return "Internal"
	}
	return fmt.Sprintf("Unknown: %d", c)
}

func (c Code) GoString() string {
	return "errcode.Code[" + c.String() + "]"
}
