package dbfs

import (
	"testing"
)

var tnodestr = []struct {
	name string
	tree node
	dest string
}{
	{"STR.1", node{name: "r", chld: nil}, "r"},
	{"STR.2", node{name: "r", chld: []*node{&node{name: "a"}}}, "r{a}"},
	{"STR.3", node{name: "r", chld: []*node{&node{name: "a"}, &node{name: "b", chld: []*node{&node{name: "c1"}, &node{name: "c2"}, &node{name: "c3"}}}, &node{name: "d"}, &node{name: "e"}}}, "r{a,b{c1,c2,c3},d,e}"},
}

var tnodeupd = []struct {
	name string
	tree node
	dest string
}{
	{"UPD.1", node{name: "r", chld: []*node{&node{name: "a"}}}, "r"},
}

var tnodewlk = []struct {
	name string
	tree node
	path []string
	dest string
	err  string
}{
	{"WLK.1", node{name: "r", chld: []*node{&node{name: "a"}}}, []string{"r"}, "r", ""},
	{"WLK.2", node{name: "r", chld: []*node{&node{name: "a"}}}, []string{"r", "a"}, "a", ""},
	{"WLK.3", node{name: "r", chld: []*node{&node{name: "a", chld: []*node{&node{name: "b"}}}}}, []string{"r", "a", "b"}, "b", ""},
	{"WLK.4", node{name: "r", chld: []*node{&node{name: "a", chld: []*node{&node{name: "b"}}}}}, []string{"r", "a", "b", ".."}, "a", ""},
	{"WLK.5", node{name: "r", chld: []*node{&node{name: "a", chld: []*node{&node{name: "b"}}}}}, []string{"r", "a", "b", "..", ".."}, "r", ""},
	{"WLK.6", node{name: "r", chld: []*node{&node{name: "a", chld: []*node{&node{name: "b"}}}}}, []string{"r", "a", "b", "..", "..", ".."}, "r", ""},
	{"WLK.7", node{name: "r", chld: []*node{&node{name: "a"}}}, []string{"r", "aa"}, "", ENOPATH},
}

func TestStrnode(t *testing.T) {
	for _, o := range tnodestr {
		s := o.tree.String()
		if s != o.dest {
			t.Errorf("%s: expected %s, got %s", o.name, o.dest, s)
		}
	}
}

func TestUpdnode(t *testing.T) {
	for _, o := range tnodeupd {
		o.tree.Updprnt(nil)

		n := o.tree.chld[0].prnt
		if n == nil {
			t.Errorf("%s: '%s'", o.name, "n is nil")
		} else if n.name != o.dest {
			t.Errorf("%s: expected %s, got %s", o.name, o.dest, n.name)
		}

	}
}

func TestWlknode(t *testing.T) {
	for _, o := range tnodewlk {
		o.tree.Updprnt(nil)

		n, err := o.tree.Wlk(o.path, nil)
		if err != nil && err.Error() != o.err {
			t.Errorf("%s: expected %s, got", o.err, err.Error())
		}
		if err != nil {
			continue
		}
		if n.name != o.dest {
			t.Errorf("%s: expected %s, got %s", o.name, o.dest, n.name)
		}

	}
}
