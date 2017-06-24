package dbfs

import (
	"9fans.net/go/plan9"
	"bytes"
	"testing"
)

var tfsproc = []struct {
	name string
	tx   plan9.Fcall
	rx   plan9.Fcall
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

	fs := Newfs()
	err = fs.proc(&b)
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

func TestFs(t *testing.T) {
	//s := "Hello"
	//buf := bytes.NewBufferString(s)
	//fmt.Fprint(buf, ", World!")
	//fmt.Println(buf.String())
	// var b bytes.Buffer // A Buffer needs no initialization.
	// b.Write([]byte("Hello "))
	// fmt.Fprintf(&b, "world!")

	for _, e := range tfsproc {
		fc, err := proc(e.tx)
		if err != nil {
			t.Errorf("%s: '%s'", e.name, err.Error())
		}
		if fc.Type == plan9.Rerror && fc.Ename != e.err {
			t.Errorf("%s: expected '%s', got '%s'", e.name, e.err, fc.Ename)
		}
		if fc.Type != e.rx.Type {
			t.Errorf("%s: expected (tx->rx): (%d->%d), got (%d->%d)", e.name, e.tx.Type, e.rx.Type, e.tx.Type, fc.Type)
		}
	}

	//m := make(map[int]int)
	//m[10]++
	//t.Fatalf("", m[5])
	//t.Fatalf("", m[10])
}
