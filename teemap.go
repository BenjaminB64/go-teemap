package teemap

import "sync"

type TeeMap[K comparable, V any] struct {
	lock *sync.RWMutex
	m map[K]V
}

// New - Make a teemap
func New[K comparable, V any]() *TeeMap[K, V] {
	return &TeeMap[K, V]{
		lock: &sync.RWMutex{},
		m: map[K]V{},
	}
}

// Set a value
func (tm *TeeMap[K, V]) Set(k K, v V) {
	tm.lock.Lock()
	defer tm.lock.Unlock()
	tm.m[k] = v
}

// Get a value
func (tm *TeeMap[K, V]) Get(k K) (v V) {
	tm.lock.RLock()
	defer tm.lock.RUnlock()
	v, _ = tm.m[k]
	return
}
