/*
Fs represents a 9P server, consumes TX, produses RX.
*/
package dbfs

import (
	"io"
	"sync"

	"9fans.net/go/plan9"
)

type fs struct {
	f      sync.Mutex
	fidref map[uint32]uint32
}

func Newfs() *fs {
	fs := &fs{
		fidref: make(map[uint32]uint32),
	}
	return fs
}

func (fs *fs) delrefs() {
	defer un(lock(&fs.f))

	for fid := range fs.fidref {
		delete(fs.fidref, fid)
	}
}

func (fs *fs) incref(fid uint32) {
	defer un(lock(&fs.f))
	fs.fidref[fid]++
}

func (fs *fs) decref(fid uint32) {
	defer un(lock(&fs.f))
	fs.fidref[fid]--
}

func (fs *fs) proc(rwc io.ReadWriteCloser) error {
	Tx, err := plan9.ReadFcall(rwc)
	if err != nil {
		return perror(EREADFCALL)
	}

	Rx := &plan9.Fcall{
		Type:   Tx.Type + 1,
		Fid:    Tx.Fid,
		Tag:    Tx.Tag,
		Newfid: Tx.Newfid,
	}

	// var f func(*plan9.Fcall, *plan9.Fcall) error
	f := fs.Bad

	switch Tx.Type {
	case plan9.Tversion:
		fs.delrefs()
		f = fs.Version
	case plan9.Tauth:
		f = fs.Auth
		//case plan9.Tattach:
		//	fs.Attach
		//case plan9.Tclunk:
		//	fs.Clunk
		//case plan9.Tflush:
		//	fs.Flush
		//case plan9.Twalk:
		//	fs.Walk
		//case plan9.Topen:
		//	fs.Open
		//case plan9.Tcreate:
		//	fs.Create
		//case plan9.Tread:
		//	fs.Read
		//case plan9.Twrite:
		//	fs.Write
		//case plan9.Tremove:
		//	fs.Remove
		//case plan9.Tstat:
		//	fs.Stat
		//case plan9.Twstat:
		//	fs.Wstat
	}
	err = f(Tx, Rx)
	if err != nil {
		Rx.Type = plan9.Rerror
		Rx.Ename = err.Error()
	}

	// send - Rx
	err = plan9.WriteFcall(rwc, Rx)
	if err != nil {
		return perror(err.Error())
	}

	return nil
}

func (fs *fs) Version(tx, rx *plan9.Fcall) error {
	if tx.Msize < plan9.IOHDRSZ {
		return perror(EVERSION)
	}
	if tx.Msize > MSIZE {
		rx.Msize = MSIZE
	} else {
		rx.Msize = tx.Msize
	}
	rx.Version = plan9.VERSION9P

	return nil
}

func (fs *fs) Auth(tx, rx *plan9.Fcall) error {
	return perror(EAUTH)
}

// oth

func (fs *fs) Bad(tx, rx *plan9.Fcall) error {
	return perror(EBAD)
}
