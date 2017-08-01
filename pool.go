package conns

import "net"

func (p *pool) Get() (net.Conn, error) {
	if conn := p.Pop(); conn != nil {
		return conn.(net.Conn), nil
	}
	return p.creater()
}

func (p *pool) Put(conn net.Conn) {
	if !p.Push(conn) {
		conn.Close()
	}
}

func (p *pool) Drain() {
	for {
		conn := p.Pop()
		if conn == nil {
			return
		}
		conn.(net.Conn).Close()
	}
}
