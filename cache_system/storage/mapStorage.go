package storage

import (
	"sync"
	"time"
)

type Node struct {
	key  string
	prev *Node
	next *Node
}

type Pair struct {
	val  string
	node *Node
	ttl  time.Time
}

type MapStorage struct {
	first    *Node
	last     *Node
	data     map[string]*Pair
	capacity int
	mu       sync.Mutex
}

func NewMapStorage(capacity int) MapStorage {
	first := Node{
		key:  "",
		prev: nil,
		next: nil,
	}
	last := Node{
		key:  "",
		prev: nil,
		next: nil,
	}
	first.prev = &last
	last.next = &first
	return MapStorage{
		first:    &first,
		last:     &last,
		data:     make(map[string]*Pair),
		capacity: capacity,
	}
}

func (m *MapStorage) Get(s string) *string {
	m.mu.Lock()
	defer m.mu.Unlock()
	pair, ok := m.data[s]
	if !ok {
		return nil
	}
	node := pair.node
	if pair.ttl.Before(time.Now()) {
		node.prev.next = node.next
		node.next.prev = node.prev
		return nil
	} else {
		m.putOnFirst(node)
		return &pair.val
	}
}

func (m *MapStorage) Put(key, val string, ttl time.Duration) {
	m.mu.Lock()
	defer m.mu.Unlock()
	pair, ok := m.data[key]
	node := &Node{}
	if ok {
		pair.val = val
		node = pair.node
		pair.ttl = time.Now().Add(ttl)
	} else {
		node.key = key
		m.data[key] = &Pair{val: val, node: node, ttl: time.Now().Add(ttl)}
	}
	m.putOnFirst(node)
	if len(m.data) > m.capacity {
		// fmt.Println("hitting capacity", m.last.next.key)
		nextNext := m.last.next.next
		delete(m.data, m.last.next.key)
		m.last.next = nextNext
		nextNext.prev = m.last
	}

}

func (m *MapStorage) putOnFirst(node *Node) {
	prev := node.prev
	next := node.next
	if prev != nil {
		prev.next = next
	}
	if next != nil {
		next.prev = prev
	}
	firstPrev := m.first.prev
	firstPrev.next = node
	node.prev = firstPrev
	node.next = m.first
	m.first.prev = node
}
