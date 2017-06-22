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

func (fs *fs) proc(rwc io.ReadWriteCloser) {
	Tx, err := plan9.ReadFcall(rwc)
	if err != nil {
		return
	}

	Rx := &plan9.Fcall{
		Type:   Tx.Type + 1,
		Fid:    Tx.Fid,
		Tag:    Tx.Tag,
		Newfid: Tx.Newfid,
	}

	switch Tx.Type {
	//case plan9.Tversion:
	//	fs.delrefs()
	//	fs.Version(Tx, Rx)
	//case plan9.Tauth:
	//	fs.Auth(Tx, Rx)
	//case plan9.Tattach:
	//	fs.Attach(Tx, Rx)
	//case plan9.Tclunk:
	//	fs.Clunk(Tx, Rx)
	//case plan9.Tflush:
	//	fs.Flush(Tx, Rx)
	//case plan9.Twalk:
	//	fs.Walk(Tx, Rx)
	//case plan9.Topen:
	//	fs.Open(Tx, Rx)
	//case plan9.Tcreate:
	//	fs.Create(Tx, Rx)
	//case plan9.Tread:
	//	fs.Read(Tx, Rx)
	//case plan9.Twrite:
	//	fs.Write(Tx, Rx)
	//case plan9.Tremove:
	//	fs.Remove(Tx, Rx)
	//case plan9.Tstat:
	//	fs.Stat(Tx, Rx)
	//case plan9.Twstat:
	//	fs.Wstat(Tx, Rx)
	default:
		fs.BadFcall(Tx, Rx)
	}

	// send - Rx
	err = plan9.WriteFcall(rwc, Rx)
	if err != nil {
		return
	}
}

func (fs *fs) Version(tx, rx *plan9.Fcall) error {
	if tx.Msize < plan9.IOHDRSZ {
		return perror("msize too small")
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
	return perror("authentication not required")
}

// oth

func (fs *fs) BadFcall(tx, rx *plan9.Fcall) error {
	return perror("bad fcall")
}
