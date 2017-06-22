package dbfs

import (
	"9fans.net/go/plan9"
	"bytes"
	"testing"
)

var txn = []string{}

var tests = []struct {
	tx string
	rx string
}{}

type rwcb struct {
	bytes.Buffer
}

func (_ *rwcb) Close() error { return nil }

func TestFs(t *testing.T) {
	//s := "Hello"
	//buf := bytes.NewBufferString(s)
	//fmt.Fprint(buf, ", World!")
	//fmt.Println(buf.String())
	// var b bytes.Buffer // A Buffer needs no initialization.
	// b.Write([]byte("Hello "))
	// fmt.Fprintf(&b, "world!")

	fs := Newfs()

	f := &plan9.Fcall{
		Type:    plan9.Tversion,
		Fid:     plan9.NOFID,
		Tag:     plan9.NOTAG,
		Msize:   131072 + plan9.IOHDRSIZE,
		Version: plan9.VERSION9P,
	}

	var b rwcb // bytes.Buffer
	err := plan9.WriteFcall(&b, f)
	if err != nil {
		t.Fatal("plan9.WriteFcall failed")
	}

	err = fs.proc(&b)
	if err != nil {
		t.Fatal(err)
	}

	//m := make(map[int]int)
	//m[10]++
	//t.Fatalf("", m[5])
	//t.Fatalf("", m[10])
}
