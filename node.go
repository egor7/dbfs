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
	// {nm, tp, fid, qid, ver, prn, chld}
	m     sync.Mutex
	nm    string
	prn   *node
	child []node
}

func Newnode(nm string) (*node, error) {
	return nil, nil
}

func (prn *node) Insnode(nm string) (*node, error) {
	return nil, nil
}

func (prn *node) Delnode(nm string) error {
	return nil
}

func (prn *node) walk(path []string) (*node, error) {
	n := prn
	for _, name := range path {
		if name == ".." {
			n = n.prn
		} else {
			found := false
			for _, nc := range n.child {
				if nc.nm == name {
					found = true
					n = &nc
				}
			}
			if found == false {
				return nil, errors.New(ENOPATH)
			}
		}
	}
	return n, nil
}

func (prn *node) String() string {

	s := prn.nm
	if prn.child != nil {
		cs := []string{}
		for _, s := range prn.child {
			cs = append(cs, s.String())
		}
		s += "{"
		s += strings.Join(cs, ",")
		s += "}"
	}
	return s
}
