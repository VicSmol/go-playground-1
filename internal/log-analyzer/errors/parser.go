// Package errors предоставляет предопределённые ошибки для парсинга логов.
package errors

import "errors"

// Parser errors for log parsing

// Common errors
var (
	ParserErrorEmptyLine = errors.New("empty line")
)

// Parser errors for JSON parser
var (
	ParserErrorInvalidJSONFormat     = errors.New("invalid JSON format")
	ParserErrorMissingFieldLevel     = errors.New("missing required field: level")
	ParserErrorMissingFieldComponent = errors.New("missing required field: component")
	ParserErrorEmptyFieldLevel       = errors.New("level field is empty")
	ParserErrorEmptyFieldComponent   = errors.New("component field is empty")
)

// Parser errors for Syslog parser
var (
	ParserErrorInvalidSyslogFormat = errors.New("invalid syslog format")
	ParserErrorUnsupportedLevel    = errors.New("unsupported log level")
)

// Parser errors for Custom parser
var (
	ParserErrorInvalidCustomFormat = errors.New("invalid custom format")
)
