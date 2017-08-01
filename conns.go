package conns

import (
	"net"

	"github.com/philchia/loop"
)

type pool struct {
	creater func() (net.Conn, error)
	loop.Loop
}

// New create a connection pool
func New(f func() (net.Conn, error), size uint) Pool {
	return &pool{
		creater: f,
		Loop:    loop.New(size),
	}
}

// Pool interface
type Pool interface {
	Get() (net.Conn, error)
	Put(net.Conn)
	Drain()
}
