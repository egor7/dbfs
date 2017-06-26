package dbfs

import (
	"testing"
)

var tnodenew = []struct {
	nm    string
	tree  node
	stree string
}{
	{"1", node{nm: "/", prn: nil, child: nil}, "/"},
	{"2", node{nm: "/", prn: nil, child: []node{node{nm: "a"}}}, "/{a}"},
	{"2", node{nm: "/", prn: nil, child: []node{node{nm: "a"}, node{nm: "b", prn: nil, child: []node{node{nm: "c"}}}}}, "/{a,b{c}}"},
}

func TestNewnode(t *testing.T) {

	for _, e := range tnodenew {

		// test nodes
		s := e.tree.String()
		if s != e.stree {
			t.Errorf("%s: expected %s, got %s", e.nm, e.stree, s)
		}
	}
}
