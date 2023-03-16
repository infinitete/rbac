package io

import (
	"errors"
	"sync"
)

var ErrRecordNotFound = errors.New("record not found")

type MemStore[T any] struct {
	mu    sync.Mutex
	items map[string]T
}

func (m *MemStore[T]) Read(key string) (T, error) {
	if record, ok := m.items[key]; ok {
		return record, nil
	}

	var t T
	return t, ErrRecordNotFound
}

func (m *MemStore[T]) Create(key string, item T) (T, error) {
	m.mu.Lock()
	m.items[key] = item
	m.mu.Unlock()

	return item, nil
}

func (m *MemStore[T]) Delete(key string) error {
	m.mu.Lock()
	delete(m.items, key)
	m.mu.Unlock()

	return nil
}
func (m *MemStore[T]) Update(key string, item T) (T, error) {
	return m.Create(key, item)
}

func GetStore[T any]() *MemStore[T] {
	return &MemStore[T]{
		mu:    sync.Mutex{},
		items: make(map[string]T),
	}
}
