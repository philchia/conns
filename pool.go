package conns

import (
	"net"
)

// this is a compiler checker that pool will implement Pool interface
var _ Pool = (*pool)(nil)

type pool struct {
	dialer func() (net.Conn, error)
	conns  chan net.Conn
}

func (p *pool) Get() (net.Conn, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		break
	}
	return p.dialer()
}

func (p *pool) Put(conn net.Conn) {
	select {
	case p.conns <- conn:
		return
	default:
		break
	}
	conn.Close()
}

func (p *pool) Drain() {
	close(p.conns)
	for conn := range p.conns {
		conn.Close()
	}
}
