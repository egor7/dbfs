package dbfs

import "testing"

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

var ttreechlds = []struct {
	name string
	tree []struct{ id, pid uint64 }
	dest []struct{ pid, cnt uint64 }
}{
	{"CHLD.1", []struct{ id, pid uint64 }{{1, 0}, {2, 0}, {3, 0}}, []struct{ pid, cnt uint64 }{{0, 3}, {1, 0}, {2, 0}, {3, 0}}},
	{"CHLD.2", []struct{ id, pid uint64 }{{1, 0}, {2, 0}, {3, 1}}, []struct{ pid, cnt uint64 }{{0, 2}, {1, 1}}},
}

func TestNewtree(t *testing.T) {
	tr := *Newtree()

	root, found := tr[ROOT]
	if !found {
		t.Errorf("Newtree: root not found")
	}

	if root.Ppath != ROOT {
		t.Errorf("Newtree: root parent is not root")
	}
}

func TestChlds(t *testing.T) {
	tr := *Newtree()

	for _, o := range ttreechlds {
		// create tree
		for _, n := range o.tree {
			tr[n.id] = &node{Name: "", Ppath: n.pid}
		}

		// compare childs counts
		for _, n := range o.dest {
			nc := len(tr.chlds(n.pid))
			if nc != int(n.cnt) {
				t.Errorf("%s: expected childs count %d, got %d", o.name, n.cnt, nc)
			}
		}

	}

}
