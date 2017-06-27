/*
Walk tree
*/
package dbfs

import "sync"

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

func (t *tree) Mkdir(path []string) error {
	// walk and create
	return nil
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

type stepFunc func(string) error

func (t *tree) walk(path []string, f stepFunc) (*node, error) {
	n := (*t)[ROOT]
	if len(path) == 0 {
		return n, nil
	}

	return n, nil
}

// dirD := newNode(fs, "d", "", "", 0775|plan9.DMDIR, 3, nil)
// fileA := newNode(fs, "fa", "", "", 0664, 4, nil)

// Qid: plan9.Qid{
//     Type: uint8(perm >> 24),
//     Vers: uint32(0),
//     Path: path,
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
