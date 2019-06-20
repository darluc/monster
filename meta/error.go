package meta

import "fmt"

// An error type for treating multiple errors as a single error.
type AggregateError []error

// Error is part of the error interface.
func (err AggregateError) Error() string {
	if len(err) == 0 {
		// This should never happen, really.
		return ""
	}
	if len(err) == 1 {
		return err[0].Error()
	}
	result := fmt.Sprintf("[%s", err[0].Error())
	for i := 1; i < len(err); i++ {
		result += fmt.Sprintf(", %s", err[i].Error())
	}
	result += "]"
	return result
}
