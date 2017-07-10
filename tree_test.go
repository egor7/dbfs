package dbfs

import "testing"

const (
	ENOROOT     = "Newtree: root not found"
	ENOROOTPRNT = "Newtree: root parent is not root"
)

type tnode struct {
	id, pid uint64
	name    string
}

type tchild struct {
	pid uint64
	cnt int
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
	dest []tchild
}{
	{"CHLD.1", []tnode{{1, 0, ""}, {2, 0, ""}, {3, 0, ""}}, []tchild{{0, 3}, {1, 0}, {2, 0}, {3, 0}}},
	{"CHLD.2", []tnode{{1, 0, ""}, {2, 0, ""}, {3, 1, ""}}, []tchild{{0, 2}, {1, 1}}},
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

var ttreemkdir = []struct {
	name string
	tree []tnode
	qid  uint64
	n    *node
	dest []tchild
	err  string
}{
	{"MKDIR.1", []tnode{{1, 0, "a"}}, 1, &node{Name: "b"}, []tchild{}, EQEXISTS},
	{"MKDIR.2", []tnode{{1, 0, "a"}}, 2, &node{Name: "b", Ppath: 2}, []tchild{}, ENOPATH},
	{"MKDIR.3", []tnode{{1, 0, "a"}}, 2, &node{Name: "b", Ppath: 1}, []tchild{{0, 1}, {1, 1}}, ""},
	{"MKDIR.4", []tnode{{1, 0, "a"}, {2, 1, "b"}}, 2, &node{Name: "b", Ppath: 1}, []tchild{}, EQEXISTS},
	{"MKDIR.5", []tnode{{1, 0, "a"}, {2, 0, "b"}}, 3, &node{Name: "b", Ppath: 0}, []tchild{}, ENEXISTS},
	{"MKDIR.6", []tnode{{1, 0, "a"}, {2, 0, "b"}}, 3, &node{Name: "c", Ppath: 0}, []tchild{{0, 3}}, ""},
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
			if n.cnt != nc {
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

		id, n, err := tr.Walk(o.path, nil)
		if err != nil {
			if o.err != err.Error() {
				t.Errorf("%s: expected %s, got %s", o.name, o.err, err.Error())
			}
			continue
		} else if o.err != "" {
			t.Errorf("%s: expected %s, got nil", o.name, o.err)
		}
		if o.dest.name != n.Name {
			t.Errorf("%s: expected %s, got %s", o.name, o.dest.name, n.Name)
		}
		if o.dest.id != id {
			t.Errorf("%s: expected %d, got %d", o.name, o.dest.id, id)
		}
	}
}

func TestMkdir(t *testing.T) {
	for _, o := range ttreemkdir {
		tr := *Newtree()
		for _, n := range o.tree {
			tr[n.id] = &node{Name: n.name, Ppath: n.pid}
		}

		err := tr.Mkdir(o.qid, o.n)
		if err != nil {
			if o.err != err.Error() {
				t.Errorf("%s: expected %s, got %s", o.name, o.err, err.Error())
			}
			continue
		} else if o.err != "" {
			t.Errorf("%s: expected %s, got nil", o.name, o.err)
		}

		for _, n := range o.dest {
			nc := len(tr.chlds(n.pid))
			if n.cnt != nc {
				t.Errorf("%s: expected childs count %d, got %d", o.name, n.cnt, nc)
			}
		}
	}
}
