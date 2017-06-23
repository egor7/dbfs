/*
Fs represents a 9P server, consumes TX, produses RX.
*/
package dbfs

import (
	"io"
	"path"
	"sync"

	"9fans.net/go/plan9"
)

type fs struct {
	f      sync.Mutex
	root   *node
	fidref map[uint32]uint32
}

func Newfs() *fs {
	fs := &fs{
		root:   &node{nm: "/", child: make(map[string]*node)},
		fidref: make(map[uint32]uint32),
	}
	fs.root.prn = fs.root
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
	case plan9.Tattach:
		f = fs.Attach
		//case plan9.Tclunk:
		//	f = fs.Clunk
		//case plan9.Tflush:
		//	f = fs.Flush
		//case plan9.Twalk:
		//	f = fs.Walk
		//case plan9.Topen:
		//	f = fs.Open
		//case plan9.Tcreate:
		//	f = fs.Create
		//case plan9.Tread:
		//	f = fs.Read
		//case plan9.Twrite:
		//	f = fs.Write
		//case plan9.Tremove:
		//	f = fs.Remove
		//case plan9.Tstat:
		//	f = fs.Stat
		//case plan9.Twstat:
		//	f = fs.Wstat
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

func (fs *fs) Attach(tx, rx *plan9.Fcall) error {
	if tx.Afid != plan9.NOFID {
		return perror(EAUTH)
	}

	aname := path.Clean(tx.Aname)
	path := split(aname)
	if len(path) == 0 {
		// TODO
	}
	// TODO
	//fs.walk(aname)

	//root, err := s.fs.Attach(tx.Uname, tx.Aname)
	//if err != nil {
	//	return err
	//}
	/// defer un(lock(&fs.f))
	//fid.node = root.node
	//fid.uid = root.uid
	//
	//stat := root.node.Stat()
	//rx.Qid = stat.Qid
	return nil
}

// oth

func (fs *fs) Bad(tx, rx *plan9.Fcall) error {
	return perror(EBAD)
}
