package dbfs

import (
	"bytes"
	"testing"

	"9fans.net/go/plan9"
)

var tsrvproc = []struct {
	nm string
	tx plan9.Fcall
	rx plan9.Fcall
	// tx   uint8
	// rx   uint8
	// fid  uint32
	// afid uint32
	err string
}{
	{"Tversion", plan9.Fcall{Type: plan9.Tversion, Fid: plan9.NOFID, Afid: plan9.NOFID}, plan9.Fcall{Type: plan9.Rversion}, ""},
	{"Tauth", plan9.Fcall{Type: plan9.Tauth, Fid: plan9.NOFID, Afid: plan9.NOFID}, plan9.Fcall{Type: plan9.Rerror}, EAUTH},
	{"Tattach.1", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID - 1}, plan9.Fcall{Type: plan9.Rerror}, EAUTH},
	{"Tattach.2", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID}, plan9.Fcall{Type: plan9.Rerror}, EEMPPATH},
	//{"Tattach.3", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID, Aname: "foo"}, plan9.Fcall{Type: plan9.Rerror}, ENOPATH},
	{"Tattach.4", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID, Aname: "/root"}, plan9.Fcall{Type: plan9.Rattach}, ""},
	{"Tattach.5", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID, Aname: "root/"}, plan9.Fcall{Type: plan9.Rattach}, ""},
	{"Tattach.6", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID, Aname: "./root"}, plan9.Fcall{Type: plan9.Rattach}, ""},
	{"Tattach.7", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID, Aname: "/root/foo/buz/../.."}, plan9.Fcall{Type: plan9.Rattach}, ""},
	{"TWalk.1", plan9.Fcall{Type: plan9.Twalk}, plan9.Fcall{Type: plan9.Rwalk}, ""},
	{"TWalk.2", plan9.Fcall{Type: plan9.Twalk, Wname: []string{"root"}}, plan9.Fcall{Type: plan9.Rwalk}, ""},
	//{"TWalk.3", plan9.Fcall{Type: plan9.Twalk, Wname: []string{"root1"}}, plan9.Fcall{Type: plan9.Rerror}, ENOPATH},
}

type rwcb struct {
	bytes.Buffer
}

func (_ *rwcb) Close() error { return nil }

func tproc(tx plan9.Fcall) (*plan9.Fcall, error) {
	tx.Tag = plan9.NOTAG
	tx.Msize = 131072 + plan9.IOHDRSIZE
	tx.Version = plan9.VERSION9P

	var b rwcb
	err := plan9.WriteFcall(&b, &tx)
	if err != nil {
		return nil, err
	}

	srv := Newsrv()

	err = srv.proc(&b)
	if err != nil {
		rx, _ := plan9.ReadFcall(&b) // WARN: 2 calls, cludge design
		return rx, err
	}

	rx, err := plan9.ReadFcall(&b)
	if err != nil {
		return rx, err
	}

	return rx, nil
}

func TestProcsrv(t *testing.T) {
	for _, o := range tsrvproc {
		n, err := tproc(o.tx)
		if err != nil {
			t.Errorf("%s: '%s'", o.nm, err.Error())
		} else if n.Type == plan9.Rerror && n.Ename != o.err {
			t.Errorf("%s: expected '%s', got '%s'", o.nm, o.err, n.Ename)
		} else if n.Type != o.rx.Type {
			t.Errorf("%s: expected (tx->rx): (%d->%d), got (%d->%d)", o.nm, o.tx.Type, o.rx.Type, o.tx.Type, n.Type)
		}
	}
}
