/*
Utilities
*/
package dbfs

import (
	"sync"
)

type perror string

func (e perror) Error() string {
	return string(e)
}

// call like: defer un(lock(&m))
func un(m *sync.Mutex) {
	m.Unlock()
}
func lock(m *sync.Mutex) *sync.Mutex {
	m.Lock()
	return m
}
