/*
Walk tree
*/
package dbfs

import (
	"errors"
	"fmt"
	"sync"
)

const (
	ROOT = iota
)

type node struct {
	m sync.Mutex

	Name string

	//Path uint64
	Vers uint32
	Type uint8

	Ppath uint64
}

type tree map[uint64]*node

func Newtree() *tree {
	t := make(map[uint64]*node, 1)

	n := node{Ppath: ROOT}
	t[ROOT] = &n

	root := tree(t)
	return &root
}

func (t *tree) chlds(n uint64) []uint64 {
	c := []uint64{}
	for p, f := range *t {
		if f.Ppath == n && f.Ppath != p {
			c = append(c, p)
		}
	}
	return c
}

func (t *tree) Mkdir(path []string) error {
	t.walk(path, func(dir string) error {
		//TODO: create
		fmt.Println("creating: ", dir)
		return nil
	})
	return nil
}

type stepFunc func(string) error

func (t *tree) walk(path []string, step stepFunc) (uint64, *node, error) {
	fmt.Println("walk")
	id := uint64(ROOT)
	n := (*t)[id]
	if len(path) == 0 {
		return id, n, nil
	}

	for _, dir := range path {
		fmt.Println("checking: ", dir)

		if step != nil {
			err := step(dir)
			if err != nil {
				return 0, nil, err
			}
		}

		if dir == ".." {
			id = n.Ppath
			n = (*t)[id]
		} else {
			found := false
			for _, cid := range t.chlds(id) {
				if (*t)[cid].Name == dir {
					found = true
					id = cid
					n = (*t)[id]
				}
			}
			if found == false {
				return 0, nil, errors.New(ENOPATH)
			}
		}
	}

	return id, n, nil
}

// dirD := newNode(fs, "d", "", "", 0775|plan9.DMDIR, 3, nil)
// fileA := newNode(fs, "fa", "", "", 0664, 4, nil)

// Qid: plan9.Qid{
//     Path: path,
//     Vers: uint32(0),
//     Type: uint8(perm >> 24),
// },

//// Set prnt links on n subtree, for manual tree construction mode
//func (n *tree) Updprnt(prnt *tree) {
//	n.prnt = prnt
//	for _, c := range n.chld {
//		c.Updprnt(n)
//	}
//}
//
//func (prnt *tree) String() string {
//
//	s := prnt.name
//	if prnt.chld != nil {
//		cs := []string{}
//		for _, s := range prnt.chld {
//			cs = append(cs, s.String())
//		}
//		s += "{"
//		s += strings.Join(cs, ",")
//		s += "}"
//	}
//	return s
//}
//
//type stepFunc func(string, *tree) error
//
//func (prnt *tree) Wlk(path []string, f stepFunc) (*tree, error) {
//	n := prnt
//	if len(path) == 0 {
//		return n, nil
//	}
//	if n.name != path[0] {
//		return nil, errors.New(ENOPATH)
//	}
//	if f != nil {
//		err := f(path[0], n)
//		if err != nil {
//			return nil, err
//		}
//	}
//	for _, name := range path[1:] {
//		if name == ".." {
//			if n.prnt != nil {
//				n = n.prnt
//			}
//		} else {
//			found := false
//			for _, nc := range n.chld {
//				if nc.name == name {
//					found = true
//					n = nc
//				}
//			}
//			if found == false {
//				return nil, errors.New(ENOPATH)
//			}
//		}
//		if f != nil {
//			err := f(name, n)
//			if err != nil {
//				return nil, err
//			}
//		}
//	}
//	return n, nil
//}
//
////func (n *tree) mkdir(path []string) {
////	_, err := n.Wlk(path, func(p string, n *tree) error {
////		wqids[i] = n.qid
////
////		return nil
////	})
////	if err != nil {
////		return err
////	}
////}
