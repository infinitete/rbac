package io

import (
	"errors"
	"sync"

	"github.com/infinitete/rbac"
)

var ErrRecordNotFound = errors.New("record not found")

type MemStore[T rbac.Storable] struct {
	mu    sync.Mutex
	items map[string]*T
}

func (m *MemStore[T]) Read(key string) (T, error) {
	if record, ok := m.items[key]; ok {
		return *record, nil
	}

	var t T
	return t, ErrRecordNotFound
}

func (m *MemStore[T]) Create(item T) error {
	m.mu.Lock()
	// TODO save
	m.items[item.Key()] = &item
	m.mu.Unlock()

	return nil
}

func (m *MemStore[T]) Delete(key string) error {
	m.mu.Lock()
	delete(m.items, key)
	m.mu.Unlock()

	return nil
}
func (m *MemStore[T]) Update(item T) error {
	return m.Create(item)
}

func GetStore[T rbac.Storable]() *MemStore[T] {
	return &MemStore[T]{
		mu:    sync.Mutex{},
		items: make(map[string]*T),
	}
}
