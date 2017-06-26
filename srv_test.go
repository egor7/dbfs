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
	{"Tattach1", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID - 1}, plan9.Fcall{Type: plan9.Rerror}, EAUTH},
	{"Tattach2", plan9.Fcall{Type: plan9.Tattach, Fid: plan9.NOFID, Afid: plan9.NOFID}, plan9.Fcall{Type: plan9.Rattach}, ""},
}

type rwcb struct {
	bytes.Buffer
}

func (_ *rwcb) Close() error { return nil }

func proc(tx plan9.Fcall) (*plan9.Fcall, error) {
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
		rx, _ := plan9.ReadFcall(&b) // !! cludge design
		return rx, err
	}

	rx, err := plan9.ReadFcall(&b)
	if err != nil {
		return rx, err
	}

	return rx, nil
}

func TestSrv(t *testing.T) {
	for _, e := range tsrvproc {
		fc, err := proc(e.tx)
		if err != nil {
			t.Errorf("%s: '%s'", e.nm, err.Error())
		} else if fc.Type == plan9.Rerror && fc.Ename != e.err {
			t.Errorf("%s: expected '%s', got '%s'", e.nm, e.err, fc.Ename)
		} else if fc.Type != e.rx.Type {
			t.Errorf("%s: expected (tx->rx): (%d->%d), got (%d->%d)", e.nm, e.tx.Type, e.rx.Type, e.tx.Type, fc.Type)
		}
	}
}
