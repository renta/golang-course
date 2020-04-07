package file

import "errors"

var (
	CopyMisconfiguration = errors.New("wrong parameter for Copy function")
	CopyFileIsDirectory  = errors.New("file is directory")
	CopyFileIsNotExists  = errors.New("file is not exists")
)
