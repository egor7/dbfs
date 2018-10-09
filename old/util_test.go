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

	for _, o := range tsplit {
		n := path.Join(split(o.src)...)

		if n != o.dst {
			t.Errorf("expected %s, got %s", o.dst, n)
		}
	}
}
