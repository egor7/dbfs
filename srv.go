/*
Plan9 file server, serve 9P Rx->Tx
*/
package dbfs

import (
	"errors"
	"io"
	"path"
	"sync"

	"9fans.net/go/plan9"
)

type srv struct {
	// {name0, tp, lsn}
	f sync.Mutex
	//tree *tree // TODO: rename

	// free the resouces according to fidref utilization
	fidref map[uint32]uint32
}

func Newsrv() *srv {
	srv := &srv{
		fidref: make(map[uint32]uint32),
	}
	//?srv.tree.prn = srv.tree
	return srv
}

func (srv *srv) delrefs() {
	defer un(lock(&srv.f))

	for fid := range srv.fidref {
		delete(srv.fidref, fid)
	}
}

func (srv *srv) incref(fid uint32) {
	defer un(lock(&srv.f))
	srv.fidref[fid]++
}

func (srv *srv) decref(fid uint32) {
	defer un(lock(&srv.f))
	srv.fidref[fid]--
}

func (srv *srv) proc(rwc io.ReadWriteCloser) error {
	Tx, err := plan9.ReadFcall(rwc)
	if err != nil {
		return errors.New(EREADFCALL)
	}

	Rx := &plan9.Fcall{
		Type:   Tx.Type + 1,
		Fid:    Tx.Fid,
		Tag:    Tx.Tag,
		Newfid: Tx.Newfid,
	}

	// var f func(*plan9.Fcall, *plan9.Fcall) error
	f := srv.Bad

	switch Tx.Type {
	case plan9.Tversion:
		srv.delrefs()
		f = srv.Version
	case plan9.Tauth:
		f = srv.Auth
	case plan9.Tattach:
		f = srv.Attach
	case plan9.Twalk:
		f = srv.Walk
		//case plan9.Tclunk:
		//	f = srv.Clunk
		//case plan9.Tflush:
		//	f = srv.Flush
		//case plan9.Topen:
		//	f = srv.Open
		//case plan9.Tcreate:
		//	f = srv.Create
		//case plan9.Tread:
		//	f = srv.Read
		//case plan9.Twrite:
		//	f = srv.Write
		//case plan9.Tremove:
		//	f = srv.Remove
		//case plan9.Tstat:
		//	f = srv.Stat
		//case plan9.Twstat:
		//	f = srv.Wstat
	}
	err = f(Tx, Rx)
	if err != nil {
		Rx.Type = plan9.Rerror
		Rx.Ename = err.Error()
	}

	// send - Rx
	err = plan9.WriteFcall(rwc, Rx)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func (srv *srv) Version(tx, rx *plan9.Fcall) error {
	if tx.Msize < plan9.IOHDRSZ {
		return errors.New(EVERSION)
	}
	if tx.Msize > MSIZE {
		rx.Msize = MSIZE
	} else {
		rx.Msize = tx.Msize
	}
	rx.Version = plan9.VERSION9P

	return nil
}

func (srv *srv) Auth(tx, rx *plan9.Fcall) error {
	return errors.New(EAUTH)
}

func (srv *srv) Attach(tx, rx *plan9.Fcall) error {
	if tx.Afid != plan9.NOFID {
		return errors.New(EAUTH)
	}

	aname := path.Clean(tx.Aname)

	path := split(aname)
	if len(path) == 0 {
		return errors.New(EEMPPATH)
	}

	// UNCOMMENT: n, err := srv.tree.Wlk(path, nil)
	// UNCOMMENT: if err != nil {
	// UNCOMMENT: 	return err
	// UNCOMMENT: }

	// WARN: ignore Uname for now
	// WARN: ignore Qid for now
	// stat := tree.srv.Stat()
	// rx.Qid = stat.Qid

	// UNCOMMENT: rx.Qid = n.qid
	return nil
}

func (srv *srv) Walk(tx, rx *plan9.Fcall) error {
	if len(tx.Wname) > plan9.MAXWELEM {
		return errors.New(EMAXPATH)
	}

	wqids := make([]plan9.Qid, len(tx.Wname))

	// UNCOMMENT: i := 0
	// UNCOMMENT: _, err := srv.tree.Wlk(tx.Wname, func(p string, n *tree) error {
	// UNCOMMENT: 	wqids[i] = n.qid
	// UNCOMMENT: 	i++
	// UNCOMMENT: 	return nil
	// UNCOMMENT: })
	// UNCOMMENT: if err != nil {
	// UNCOMMENT: 	return err
	// UNCOMMENT: }

	// WARN: ignore Uname for now
	// WARN: ignore Qid for now
	// stat := tree.srv.Stat()
	// rx.Qid = stat.Qid

	rx.Wqid = wqids
	return nil
}

// oth: srv:fs:fid

func (srv *srv) Bad(tx, rx *plan9.Fcall) error {
	return errors.New(EBAD)
}
