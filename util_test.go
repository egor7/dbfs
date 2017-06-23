package dbfs

import (
	"path"
	"testing"
)

var tsplit = []struct {
	src string
	dst string
}{
	{"/tables/t1", "tables/t1"},
	{"packages/p1", "packages/p1"},
	{"d1/..", "."},
	{"d1/../d2", "d2"},
	{"d1/../../d2/", "../d2"},
}

func TestUtil(t *testing.T) {

	for _, e := range tsplit {
		dst := path.Join(split(e.src)...)

		if dst != e.dst {
			t.Errorf("expected %s, got %s", e.dst, dst)
		}
	}
}
