/*
Walk tree
*/
package dbfs

import (
	"errors"
	"strings"
	"sync"

	"9fans.net/go/plan9"
)

type node struct {
	// {name, tp, fid, qid, ver, prnt, chld}
	m sync.Mutex

	name string
	prnt *node
	chld []*node

	qid plan9.Qid
}

type stepFunc func(string, *node) error

func (prnt *node) Ins(name string) (*node, error) {
	// TODO: make
	return nil, nil
}

func (prnt *node) Del(name string) error {
	// TODO: make
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

func (prnt *node) Wlk(path []string, f stepFunc) (*node, error) {
	n := prnt
	if n.name != path[0] {
		return nil, errors.New(ENOPATH)
	}
	if f != nil {
		err := f(path[0], n)
		if err != nil {
			return nil, err
		}
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
		if f != nil {
			err := f(name, n)
			if err != nil {
				return nil, err
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
