package dbfs

import "testing"

const (
	ENOROOT     = "Newtree: root not found"
	ENOROOTPRNT = "Newtree: root parent is not root"
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

type tnode struct {
	id, pid uint64
	name    string
}

var ttreenew = []struct {
	name string
	root tnode
}{
	{"NEW.1", tnode{ROOT, ROOT, ""}},
}

var ttreechlds = []struct {
	name string
	tree []tnode
	dest []struct{ pid, cnt uint64 }
}{
	{"CHLD.1", []tnode{{1, 0, ""}, {2, 0, ""}, {3, 0, ""}}, []struct{ pid, cnt uint64 }{{0, 3}, {1, 0}, {2, 0}, {3, 0}}},
	{"CHLD.2", []tnode{{1, 0, ""}, {2, 0, ""}, {3, 1, ""}}, []struct{ pid, cnt uint64 }{{0, 2}, {1, 1}}},
}

var ttreewalk = []struct {
	name string
	tree []tnode
	path []string
	dest tnode
	err  string
}{
	{"WALK.1", []tnode{{1, 0, "a"}}, []string{}, tnode{ROOT, 0, ""}, ""},
	{"WALK.2", []tnode{{1, 0, "a"}}, []string{".."}, tnode{ROOT, 0, ""}, ""},
	{"WALK.3", []tnode{{1, 0, "a"}}, []string{"..", ".."}, tnode{ROOT, 0, ""}, ""},
	{"WALK.4", []tnode{{1, 0, "a"}}, []string{"a", "..", ".."}, tnode{ROOT, 0, ""}, ""},
	{"WALK.5", []tnode{{1, 0, "a"}}, []string{"a"}, tnode{1, 0, "a"}, ""},
	{"WALK.6", []tnode{{1, 0, "a"}, {2, 1, "b"}}, []string{"a", "b"}, tnode{2, 0, "b"}, ""},
	{"WALK.7", []tnode{{1, 0, "a"}, {2, 1, "b"}}, []string{"a", "b", ".."}, tnode{1, 0, "a"}, ""},
	{"WALK.8", []tnode{{1, 0, "a"}, {2, 1, "b"}}, []string{"a", "b", "..", ".."}, tnode{ROOT, 0, ""}, ""},
	{"WALK.9", []tnode{{1, 0, "a"}}, []string{"b"}, tnode{}, ENOPATH},
	{"WALK.10", []tnode{{1, 0, "a"}}, []string{"a", "b"}, tnode{}, ENOPATH},
	{"WALK.11", []tnode{{1, 0, "a"}}, []string{"a", "b", ".."}, tnode{}, ENOPATH},
}

func TestNewtree(t *testing.T) {
	for _, o := range ttreenew {
		tr := *Newtree()
		root, found := tr[o.root.id]
		if !found {
			t.Errorf(ENOROOT)
		} else if root.Ppath != o.root.pid {
			t.Errorf(ENOROOTPRNT)
		}
	}
}

func TestChlds(t *testing.T) {
	for _, o := range ttreechlds {
		tr := *Newtree()
		for _, n := range o.tree {
			tr[n.id] = &node{Name: n.name, Ppath: n.pid}
		}

		for _, n := range o.dest {
			nc := len(tr.chlds(n.pid))
			if nc != int(n.cnt) {
				t.Errorf("%s: expected childs count %d, got %d", o.name, n.cnt, nc)
			}
		}

	}
}

func TestWalk(t *testing.T) {
	for _, o := range ttreewalk {
		tr := *Newtree()
		for _, n := range o.tree {
			tr[n.id] = &node{Name: n.name, Ppath: n.pid}
		}

		id, n, err := tr.walk(o.path, nil)
		if err != nil {
			if o.err != err.Error() {
				t.Errorf("%s: expected %s, got %s", o.name, o.err, err.Error())
			}
			continue
		}
		if o.dest.name != n.Name {
			t.Errorf("%s: expected %s, got %s", o.name, o.dest.name, n.Name)
		}
		if o.dest.id != id {
			t.Errorf("%s: expected %d, got %d", o.name, o.dest.id, id)
		}
	}
}
