package dbfs

import (
	"9fans.net/go/plan9"
	"bytes"
	"testing"
)

var tests = []struct {
	name string
	tx   uint8
	rx   uint8
	fid  uint32
	afid uint32
	err  string
}{
	{"Tversion", plan9.Tversion, plan9.Rversion, plan9.NOFID, plan9.NOFID, ""},
	{"Tauth", plan9.Tauth, plan9.Rerror, plan9.NOFID, plan9.NOFID, EAUTH},
	{"Tattach1", plan9.Tattach, plan9.Rerror, plan9.NOFID, plan9.NOFID - 1, EAUTH},
	{"Tattach2", plan9.Tattach, plan9.Rattach, plan9.NOFID, plan9.NOFID, ""},
}

type rwcb struct {
	bytes.Buffer
}

func (_ *rwcb) Close() error { return nil }

func proc(tp uint8, fid uint32, afid uint32) (*plan9.Fcall, error) {
	f := &plan9.Fcall{
		Type:    tp,
		Fid:     fid,
		Afid:    afid,
		Tag:     plan9.NOTAG,
		Msize:   131072 + plan9.IOHDRSIZE,
		Version: plan9.VERSION9P,
	}

	var b rwcb
	err := plan9.WriteFcall(&b, f)
	if err != nil {
		return f, err
	}

	fs := Newfs()
	err = fs.proc(&b)
	if err != nil {
		f, _ = plan9.ReadFcall(&b) // !! cludge design
		return f, err
	}

	f, err = plan9.ReadFcall(&b)
	if err != nil {
		return f, err
	}

	return f, nil
}

func TestFs(t *testing.T) {
	//s := "Hello"
	//buf := bytes.NewBufferString(s)
	//fmt.Fprint(buf, ", World!")
	//fmt.Println(buf.String())
	// var b bytes.Buffer // A Buffer needs no initialization.
	// b.Write([]byte("Hello "))
	// fmt.Fprintf(&b, "world!")

	for _, e := range tests {
		fc, err := proc(e.tx, e.fid, e.afid)

		if err != nil {
			t.Errorf("%s: '%s'", e.name, err.Error())
		}
		if fc.Type == plan9.Rerror && fc.Ename != e.err {
			t.Errorf("%s: expected '%s', got '%s'", e.name, e.err, fc.Ename)
		}
		if fc.Type != e.rx {
			t.Errorf("%s: expected (tx->rx): (%d->%d), got (%d->%d)", e.name, e.tx, e.rx, e.tx, fc.Type)
		}
	}

	//m := make(map[int]int)
	//m[10]++
	//t.Fatalf("", m[5])
	//t.Fatalf("", m[10])
}
