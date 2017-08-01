package conns

import (
	"net"

	"github.com/philchia/loop"
)

// New create a connection pool
func New(dialer func() (net.Conn, error), size uint) Pool {
	return &pool{
		dialer: dialer,
		loop:   loop.New(size),
	}
}

// Pool interface
type Pool interface {
	Get() (net.Conn, error)
	Put(net.Conn)
	Drain()
}
