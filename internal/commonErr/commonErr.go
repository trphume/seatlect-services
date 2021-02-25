package commonErr

import "errors"

var (
	INTERNAL  = errors.New("INTERNAL ERROR")
	DUPLICATE = errors.New("DUPLICATE")
	NOTFOUND  = errors.New("NOT FOUND")
	CONFLICT  = errors.New("CONFLICT")
	INVALID   = errors.New("INVALID")
)
