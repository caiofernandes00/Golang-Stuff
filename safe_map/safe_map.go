package safe_map

import (
	"fmt"
	"sync"
)

type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

func New[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V),
	}
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mu.Lock()

	m.data[key] = value

	m.mu.Unlock()
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mu.RLock()

	value, ok := m.data[key]
	if !ok {
		return value, fmt.Errorf("key %v not found", key)
	}

	m.mu.RLock()
	return value, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mu.Lock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	m.data[key] = value

	m.mu.Unlock()
	return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
	m.mu.Lock()

	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	delete(m.data, key)

	m.mu.Unlock()
	return nil
}

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.mu.RLock()

	_, ok := m.data[key]
	if ok {
		return true
	}

	m.mu.RUnlock()
	return false
}
