package ierror

import "fmt"

type IError struct {
	Info       string
	Additional string
}

func (e *IError) Error() string {
	if e.Additional != "" {
		return fmt.Sprintf("%s: %s", e.Info, e.Additional)
	}
	return fmt.Sprintf("%s", e.Info)
}

func New(info string) IError {
	return IError{
		Info:       info,
		Additional: "",
	}
}

// Extended - exntends existing error with additional info
func Extended(error *IError, additional string) IError {
	return IError{
		Info:       error.Info,
		Additional: additional,
	}
}

// NewExtended - returns `IError` with `info` & `additional` info
func NewExtended(info string, additional string) IError {
	return IError{
		Info:       info,
		Additional: additional,
	}
}
