package loop

// Loop represent a loop queue
type Loop interface {
	Push(interface{}) bool
	Pop() interface{}
}

// loop queue
type loop struct {
	// read/write position
	read, write uint
	// internal data holder
	data []interface{}
	// capicity
	cap uint
	len uint
}

// New create a Loop instance
func New(cap uint) Loop {
	return &loop{
		data: make([]interface{}, cap),
		cap:  cap,
	}
}

// Push a element into loop, return true if success, else return false
func (l *loop) Push(i interface{}) bool {
	if l.full() {
		return false
	}
	l.data[l.write] = i
	l.write = (l.write + 1) % l.cap
	l.len++
	return true
}

// Pop a element from loop, return nil if empty
func (l *loop) Pop() interface{} {
	if l.empty() {
		return nil
	}
	i := l.data[l.read]
	l.read = (l.read + 1) % l.cap
	l.len--
	return i
}

func (l *loop) empty() bool {
	return l.len == 0
}

func (l *loop) full() bool {
	return l.len == l.cap
}
