package dbfs

import (
	"net"
	// "9fans.net/go/plan9"
)

// handle network connections
func Listen(network, addr string) error {
	l, err := net.Listen(network, addr)
	if err != nil {
		return err
	}

	for {
		rwc, err := l.Accept()
		if err != nil {
			continue
		}

		go func(rwc net.Conn) {
			defer rwc.Close()

			//fs := &fs{
			//	fidref: make(map[uint32]uint32),
			//}

			//proc(rwc)
		}(rwc)
	}
}
