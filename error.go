package semver

import "fmt"

type invalidNumericError struct {
	part  string
	value string
}

func (e *invalidNumericError) Error() string {
	return fmt.Sprintf("invalid numeric identifier (leading zeros): %s = %s", e.part, e.value)
}

func newInvalidNumericError(part, value string) *invalidNumericError {
	return &invalidNumericError{
		part:  part,
		value: value,
	}
}

// ParseError represents an error caused during parsing.
type ParseError struct {
	message string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", e.message)
}

func newParseError(msg string) *ParseError {
	return &ParseError{message: msg}
}
