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
	EEMPPATH = "E.srv: Path empty"
	EMAXPATH = "E.srv: Too many names in walk"

	ENOPATH  = "E.tree: Path does not exist"
	EQEXISTS = "E.tree: Path Qid already exist"
	ENEXISTS = "E.tree: Path Name already exist"
)
