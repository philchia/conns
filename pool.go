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

// Get fetch a conn from pool, create one if empty
func (p *pool) Get() (net.Conn, error) {
	select {
	case conn := <-p.conns:
		return conn, nil
	default:
		break
	}
	return p.dialer()
}

// Put conn back to pool, close if full
func (p *pool) Put(conn net.Conn) {
	select {
	case p.conns <- conn:
		return
	default:
		break
	}
	conn.Close()
}

// Drain the pool, close all conns
func (p *pool) Drain() {
	close(p.conns)
	for conn := range p.conns {
		conn.Close()
	}
}
