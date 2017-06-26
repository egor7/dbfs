/*
Walk tree
*/
package dbfs

import (
	"errors"
	"strings"
	"sync"
)

type node struct {
	// {name, tp, fid, qid, ver, prnt, chld}
	m    sync.Mutex
	name string
	prnt *node
	chld []*node
}

func Newnode(name string) (*node, error) {
	return nil, nil
}

func (prnt *node) Insnode(name string) (*node, error) {
	return nil, nil
}

func (prnt *node) Delnode(name string) error {
	return nil
}

// Set prnt links on n subtree
func (n *node) Updprnt(prnt *node) {
	//fmt.Println(&n)
	n.prnt = prnt
	for _, c := range n.chld {
		c.Updprnt(n)
	}
}

func (prnt *node) Wlknode(path []string) (*node, error) {
	n := prnt
	if n.name != path[0] {
		return nil, errors.New(ENOPATH)
	}
	for _, name := range path[1:] {
		if name == ".." {
			if n.prnt != nil {
				n = n.prnt
			}
		} else {
			found := false
			for _, nc := range n.chld {
				if nc.name == name {
					found = true
					n = nc
				}
			}
			if found == false {
				return nil, errors.New(ENOPATH)
			}
		}
	}
	return n, nil
}

func (prnt *node) String() string {

	s := prnt.name
	if prnt.chld != nil {
		cs := []string{}
		for _, s := range prnt.chld {
			cs = append(cs, s.String())
		}
		s += "{"
		s += strings.Join(cs, ",")
		s += "}"
	}
	return s
}
