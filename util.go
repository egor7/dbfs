/*
Utilities
*/
package dbfs

import (
	"sync"
)

func lock(m *sync.Mutex) *sync.Mutex {
	m.Lock()
	return m
}

func un(m *sync.Mutex) {
	m.Unlock()
}
