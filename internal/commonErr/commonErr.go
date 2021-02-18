package commonErr

import "errors"

var (
	INTERNAL  = errors.New("iINTERNAL ERROR")
	DUPLICATE = errors.New("DUPLICATE")
	NOTFOUND  = errors.New("NOT FOUND")
	CONFLICT  = errors.New("CONFLICT")
	INVALID   = errors.New("INVALID")
)
