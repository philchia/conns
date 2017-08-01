package conns

import "net"

func (p *pool) Get() (net.Conn, error) {
	if conn := p.loop.Pop(); conn != nil {
		return conn.(net.Conn), nil
	}
	return p.dialer()
}

func (p *pool) Put(conn net.Conn) {
	if !p.loop.Push(conn) {
		conn.Close()
	}
}

func (p *pool) Drain() {
	for {
		conn := p.loop.Pop()
		if conn == nil {
			return
		}
		conn.(net.Conn).Close()
	}
}
