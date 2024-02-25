package transx

import "errors"

var (
	ErrSrcNotStruct    = errors.New("source is not a struct")
	ErrDstNotPtrStruct = errors.New("destination is not pointer of a struct")
)
