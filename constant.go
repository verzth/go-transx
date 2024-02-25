package transx

import "errors"

var (
	ErrSrcNotStruct    = errors.New("source is not a struct")
	ErrDstNotPtrStruct = errors.New("destination is not pointer of a struct")
	ErrSrcSliceNil     = errors.New("source slice is nil")
	ErrSrcNotSlice     = errors.New("source is not a slice")
	ErrDstSliceNil     = errors.New("destination slice is nil")
	ErrDstNotSlice     = errors.New("destination is not a slice")
	ErrDstSliceNotPtr  = errors.New("destination is not pointer of a slice")
)
