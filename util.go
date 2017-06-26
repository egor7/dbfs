/*
Utilities
*/
package dbfs

import (
	"strings"
	"sync"
)

// call like: defer un(lock(&m))
func un(m *sync.Mutex) {
	m.Unlock()
}
func lock(m *sync.Mutex) *sync.Mutex {
	m.Lock()
	return m
}

func split(path string) []string {
	if len(path) == 0 || path == "/" || path == "." {
		return []string{}
	}

	if len(path) > 0 && path[0] == '/' {
		path = path[1:]
	}
	return strings.Split(path, "/")
}
