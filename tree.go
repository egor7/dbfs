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

	//Qpath uint64
	Qvers uint32
	Qtype uint8

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

func (t *tree) chlds(qpath uint64) []uint64 {
	c := []uint64{}
	for p, f := range *t {
		if f.Ppath == qpath && f.Ppath != p {
			c = append(c, p)
		}
	}
	return c
}

func (t *tree) Mkdir(qpath uint64, n *node) error {
	if _, found := (*t)[qpath]; found {
		return errors.New(EQEXISTS)
	}

	if _, found := (*t)[n.Ppath]; !found {
		return errors.New(ENOPATH)
	}

	//fmt.Println(qpath)
	found := false
	for _, qc := range t.chlds(n.Ppath) {
		//fmt.Println((*t)[qc].Name, " -- ", n.Name)
		if (*t)[qc].Name == n.Name {
			found = true
		}
	}
	if found == true {
		return errors.New(ENEXISTS)
	}

	(*t)[qpath] = n

	return nil
}

type stepFunc func(string) error

func (t *tree) Walk(path []string, step stepFunc) (uint64, *node, error) {
	//fmt.Println("walk")
	qpath := uint64(ROOT)
	n := (*t)[qpath]
	if len(path) == 0 {
		return qpath, n, nil
	}

	for _, dir := range path {
		//fmt.Println("checking: ", dir)

		if dir == ".." {
			qpath = n.Ppath
			n = (*t)[qpath]
		} else {
			found := false
			for _, qc := range t.chlds(qpath) {
				if (*t)[qc].Name == dir {
					found = true
					qpath = qc
					n = (*t)[qpath]
				}
			}
			if found == false {
				return 0, nil, errors.New(ENOPATH)
			}
		}

		if step != nil {
			err := step(dir)
			if err != nil {
				return 0, nil, err
			}
		}
	}

	return qpath, n, nil
}

// dirD := newNode(fs, "d", "", "", 0775|plan9.DMDIR, 3, nil)
// fileA := newNode(fs, "fa", "", "", 0664, 4, nil)

// Qid: plan9.Qid{
//     Path: path,
//     Vers: uint32(0),
//     Type: uint8(perm >> 24),
// },

func (t *tree) String() string {
	s := ""
	for qpath, n := range *t {
		s += fmt.Sprintf("%d->%d:{'%s'}; ", qpath, n.Ppath, n.Name)
	}
	return s
}
