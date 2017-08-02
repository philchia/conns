package conns

import (
	"net"
)

// New create a connection pool
func New(dialer func() (net.Conn, error), size uint) Pool {
	return &pool{
		dialer: dialer,
		conns:  make(chan net.Conn, size),
	}
}

// Pool interface
type Pool interface {
	Get() (net.Conn, error)
	Put(net.Conn)
	Drain()
}
