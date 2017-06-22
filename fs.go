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
	// recv - Tx
	Tx, err := plan9.ReadFcall(rwc)
	if err != nil {
		return
	}

	// proc - Tx
	switch Tx.Type {
	case plan9.Tversion:
		fs.delrefs()
	case plan9.Tauth:
		// nothing
	case plan9.Twalk:
		fs.incref(Tx.Fid)
		// TODO Tx.Newfid exists too, increase?
	default:
		fs.incref(Tx.Fid)
	}

	// proc - Rx
	Rx := &plan9.Fcall{
		Type:   Tx.Type + 1,
		Fid:    Tx.Fid,
		Tag:    Tx.Tag,
		Newfid: Tx.Newfid,
	}

	// [TODO] make work

	switch Rx.Type {
	case plan9.Rversion, plan9.Rauth:
		// nothing
	//	case plan9.Rattach:
	//		c.f.Lock()
	//		c.uid = req.Fid.uid
	//		c.f.Unlock()
	//		req.Fid.decRef()
	//		c.DelFid(req.Fid.num)
	case plan9.Rwalk, plan9.Rclunk:
		fs.decref(Tx.Fid)
		// delfid?
	case plan9.Rerror:
		fs.decref(Tx.Fid)
	default:
		fs.decref(Rx.Fid)
	}

	// send - Rx
	err = plan9.WriteFcall(rwc, Rx)
	if err != nil {
		return
	}
}
