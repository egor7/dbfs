/*
Walk tree
*/
package dbfs

import (
	"sync"

	"9fans.net/go/plan9"
)

type node struct {
	// {name, tp, fid, qid, ver, prnt, chld}
	m sync.Mutex

	name string

	pqid plan9.Qid
}

type tree map[plan9.Qid]*node

func Newtree() *tree {
	t := make(map[plan9.Qid]*node, 1)

	q := plan9.Qid{Path: 0, Vers: 0, Type: 0}
	n := node{pqid: q}
	t[q] = &n

	root := tree(t)
	return &root
}

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
