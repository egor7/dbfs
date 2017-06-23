/*
Db maps nodes-tree with database queries results.
*/
package dbfs

import (
	"sync"
)

type node struct {
	m     sync.Mutex
	nm    string
	prn   *node
	child map[string]*node
}

func (n *node) walk(path []string) (*node, error) {
	p := n
	for _, name := range path {
		if name == ".." {
			p = p.prn
		} else {
			ch, found := p.child[name]
			if found {
				p = ch
			} else {
				return p, perror(ENOPATH)
			}
		}
	}
	return p, nil
}
