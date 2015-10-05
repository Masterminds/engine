package form

import (
	"errors"
	"sync"
	"time"
)

// SweepInterval is a suggested interval for sweeping the cache.
//
// It is not mandatory that caching backends use SweepInterval, but the
// default caching mechanism uses it to indicate how frequently it should
// purge the cache of expired records.
var SweepInterval = 5 * time.Minute

// FormNotFound indicates that a form is not in the cache.
//
// Expired records should also return this error.
var ErrFormNotFound = errors.New("Form not found")

// CacheVal describes a value in the cache.
//
// Timeout indicates when the current form is no longer valid, and
// can be cleaned up.
type CacheVal struct {
	exp  time.Time
	form *Form
}

// Cache provides storage for forms.
//
// Generally, a cache is not used directly. Instead, the FormHandler is
// used.
//
// Cache implementations are required to handle expiration internally.
type Cache interface {
	Get(id string) (*Form, error)
	Set(id string, f *Form, expires time.Time) error
	Remove(id string) error
}

// NewCache returns a new Cache backed by an in-memory cache.
//
// TODO: An in-memory cache is currently designed to last the lifetime of
// an application. There is no way to stop and cleanup the cache.
func NewCache() Cache {
	mc := &memoryCache{
		store:  map[string]*CacheVal{},
		ticker: time.NewTicker(SweepInterval),
	}
	go mc.purge()
	return mc
}

// memoryCache stores forms in memory.
type memoryCache struct {
	mx     sync.RWMutex
	store  map[string]*CacheVal
	ticker *time.Ticker
}

func (m *memoryCache) purge() {
	// Run a simple mark-and-sweep every
	// tick.
	for now := range m.ticker.C {
		mark := []string{}
		m.mx.Lock()
		for k, v := range m.store {
			if now.After(v.exp) {
				mark = append(mark, k)
			}
		}
		for _, id := range mark {
			delete(m.store, id)
		}
		m.mx.Unlock()
	}
}

func (m *memoryCache) Get(id string) (*Form, error) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.store[id]
	if !ok {
		return nil, ErrFormNotFound
	}
	// Expire an entry if necessary.
	if time.Now().After(val.exp) {
		delete(m.store, id)
		return nil, ErrFormNotFound
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
