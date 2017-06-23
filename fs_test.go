package dbfs

import (
	"9fans.net/go/plan9"
	"bytes"
	"testing"
)

var txn = []string{}

var tests = []struct {
	name string
	tx   uint8
	rx   uint8
	err  string
}{
	{"Tversion", plan9.Tversion, plan9.Rversion, ""},
	{"Tauth", plan9.Tauth, plan9.Rerror, EAUTH},
}

type rwcb struct {
	bytes.Buffer
}

func (_ *rwcb) Close() error { return nil }

func fc(tx uint8) (*plan9.Fcall, error) {
	f := &plan9.Fcall{
		Type:    tx,
		Fid:     plan9.NOFID,
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
		f, _ = plan9.ReadFcall(&b)
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

	for _, f := range tests {
		rx, err := fc(f.tx)

		if err != nil {
			t.Errorf("%s: '%s'", f.name, err.Error())
		}
		if rx.Type == plan9.Rerror && rx.Ename != f.err {
			t.Errorf("%s: expected '%s', got '%s'", f.name, f.err, rx.Ename)
		}
		if rx.Type != f.rx {
			t.Errorf("%s: expected (tx->rx): (%d->%d), got (%d->%d)", f.name, f.tx, f.rx, f.tx, rx.Type)
		}
	}

	//m := make(map[int]int)
	//m[10]++
	//t.Fatalf("", m[5])
	//t.Fatalf("", m[10])
}
