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
	EREADFCALL = "E.srv ReadFcall goes wrong"

	EVERSION = "E.srv: msize too small"
	EAUTH    = "E.srv: authentication not required"
	EBAD     = "E.srv: bad fcall"

	ENOPATH = "E.node: File does not exist"
)
