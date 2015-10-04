package form

import (
	"errors"
	"sync"
	"time"
)

var FormNotFound = errors.New("Form not found")

type Cache interface {
	Get(id string) (*Form, error)
	Set(id string, f *Form, expires time.Time) error
	Remove(id string) error
}

// New returns a new Cache backed by an in-memory cache.
func NewCache() Cache {
	return &memoryCache{
		store: map[string]*CacheVal{},
	}
}

// memoryCache stores forms in memory.
type memoryCache struct {
	mx    sync.RWMutex
	store map[string]*CacheVal
}

// CacheVal describes a value in the cache.
//
// Timeout indicates when the current form is no longer valid, and
// can be cleaned up.
type CacheVal struct {
	exp  time.Time
	form *Form
}

func (m *memoryCache) Get(id string) (*Form, error) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.store[id]
	if !ok {
		return nil, FormNotFound
	}
	return val.form, nil
}

func (m *memoryCache) Set(id string, f *Form, expires time.Time) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.store[id] = &CacheVal{expires, f}
	return nil
}

func (m *memoryCache) Remove(id string) error {
	m.mx.Lock()
	defer m.mx.Unlock()
	delete(m.store, id)
	return nil
}
