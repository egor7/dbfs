package dbfs

import (
	"testing"

	"9fans.net/go/plan9"
)

//import (
//	"testing"
//)
//
//var ttreestr = []struct {
//	name string
//	tree tree
//	dest string
//}{
//	{"STR.1", tree{name: "r", chld: nil}, "r"},
//	{"STR.2", tree{name: "r", chld: []*tree{&tree{name: "a"}}}, "r{a}"},
//	{"STR.3", tree{name: "r", chld: []*tree{&tree{name: "a"}, &tree{name: "b", chld: []*tree{&tree{name: "c1"}, &tree{name: "c2"}, &tree{name: "c3"}}}, &tree{name: "d"}, &tree{name: "e"}}}, "r{a,b{c1,c2,c3},d,e}"},
//}
//
//var ttreeupd = []struct {
//	name string
//	tree tree
//	dest string
//}{
//	{"UPD.1", tree{name: "r", chld: []*tree{&tree{name: "a"}}}, "r"},
//}
//
//var ttreewlk = []struct {
//	name string
//	tree tree
//	path []string
//	dest string
//	err  string
//}{
//	{"WLK.1", tree{name: "r", chld: []*tree{&tree{name: "a"}}}, []string{"r"}, "r", ""},
//	{"WLK.2", tree{name: "r", chld: []*tree{&tree{name: "a"}}}, []string{"r", "a"}, "a", ""},
//	{"WLK.3", tree{name: "r", chld: []*tree{&tree{name: "a", chld: []*tree{&tree{name: "b"}}}}}, []string{"r", "a", "b"}, "b", ""},
//	{"WLK.4", tree{name: "r", chld: []*tree{&tree{name: "a", chld: []*tree{&tree{name: "b"}}}}}, []string{"r", "a", "b", ".."}, "a", ""},
//	{"WLK.5", tree{name: "r", chld: []*tree{&tree{name: "a", chld: []*tree{&tree{name: "b"}}}}}, []string{"r", "a", "b", "..", ".."}, "r", ""},
//	{"WLK.6", tree{name: "r", chld: []*tree{&tree{name: "a", chld: []*tree{&tree{name: "b"}}}}}, []string{"r", "a", "b", "..", "..", ".."}, "r", ""},
//	{"WLK.7", tree{name: "r", chld: []*tree{&tree{name: "a"}}}, []string{"r", "aa"}, "", ENOPATH},
//}
//
//func TestStrtree(t *testing.T) {
//	for _, o := range ttreestr {
//		s := o.tree.String()
//		if s != o.dest {
//			t.Errorf("%s: expected %s, got %s", o.name, o.dest, s)
//		}
//	}
//}
//
//func TestUpdtree(t *testing.T) {
//	for _, o := range ttreeupd {
//		o.tree.Updprnt(nil)
//
//		n := o.tree.chld[0].prnt
//		if n == nil {
//			t.Errorf("%s: '%s'", o.name, "n is nil")
//		} else if n.name != o.dest {
//			t.Errorf("%s: expected %s, got %s", o.name, o.dest, n.name)
//		}
//
//	}
//}
//
//func TestWlktree(t *testing.T) {
//	for _, o := range ttreewlk {
//		o.tree.Updprnt(nil)
//
//		n, err := o.tree.Wlk(o.path, nil)
//		if err != nil && err.Error() != o.err {
//			t.Errorf("%s: expected %s, got", o.err, err.Error())
//		}
//		if err != nil {
//			continue
//		}
//		if n.name != o.dest {
//			t.Errorf("%s: expected %s, got %s", o.name, o.dest, n.name)
//		}
//
//	}
//}

func TestNewtree(t *testing.T) {
	n := Newtree()
	q0 := plan9.Qid{Path: 0, Vers: 0, Type: 0}

	root, found := (*n)[q0]
	if !found {
		t.Errorf("Newtree: root not found")
	}

	if root.pqid != q0 {
		t.Errorf("Newtree: root parent is not root")
	}
}
