/*
Constants
*/
package dbfs

import (
	"9fans.net/go/plan9"
)

const (
	MSIZE  = 128*1024 + plan9.IOHDRSZ
	IOUNIT = 128 * 1024
)

const (
	EREADFCALL = "E ReadFcall goes wrong"

	EVERSION = "E msize too small"
	EAUTH    = "E authentication not required"
	EBAD     = "E bad fcall"

	ENOPATH = "E file does not exist"
)
