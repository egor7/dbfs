/*
Fs represents a 9P server, consumes TX, produses RX.
*/
package dbfs

import (
//	"io"
//	"9fans.net/go/plan9"
)

/*
func proc(rwc io.ReadWriteCloser) {
	// recv - Tx
	Tx, err := plan9.ReadFcall(rwc)
	if err != nil {
		return
	}




	// proc - Tx
	switch Tx.Type {
	case plan9.Tversion:
		c.f.Lock() // abort all outstanding I/O
		for num := range c.fidmap {
			delete(c.fidmap, num)
		}
		c.f.Unlock()
	case plan9.Tauth:
		// nothing
	default:
		req.Fid = c.GetFid(Tx.Fid)
		req.Fid.incRef()
		if Tx.Type == plan9.Twalk {
			req.Fid.New = c.GetFid(Tx.Newfid)
		}
	}


	// proc - Rx
	Rx := &plan9.Fcall{}

	if req.Err != nil {
		Rx.Type = plan9.Rerror
		Rx.Ename = req.Err.Error()
	} else {
		Rx.Type = Tx.Type + 1
		Rx.Fid = Tx.Fid
	}
	Rx.Tag = Tx.Tag


	switch Rx.Type {
	case plan9.Rversion, plan9.Rauth:
		// nothing
	case plan9.Rattach:
		c.f.Lock()
		c.uid = req.Fid.uid
		c.f.Unlock()
		req.Fid.decRef()
		c.DelFid(req.Fid.num)
	case plan9.Rwalk, plan9.Rclunk:
		req.Fid.decRef()
		c.DelFid(req.Fid.num)
	case plan9.Rerror:
		if req.Fid != nil {
			req.Fid.decRef()
		}
	default:
		req.Fid.decRef()
	}

	if c.getErr() == nil {
		reqout <- req
	}






	// send - Rx
	err = plan9.WriteFcall(rwc, Rx)
	if err != nil {
		return
	}
}
*/
