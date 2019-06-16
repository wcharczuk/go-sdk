package cache

import "sync"

// Cache is a type that implements the cache interface.
type Cache interface {
	Has(key interface{}) bool
	GetOrSet(key interface{}, valueProvider func() interface{}, options ...ValueOption) (interface{}, bool)
	Set(key, value interface{}, options ...ValueOption)
	Get(key interface{}) (interface{}, bool)
	Remove(key interface{}) (interface{}, bool)
}

// Locker is a cache type that supports external control of locking for both exclusive and reader/writer locks.
type Locker interface {
	sync.Locker
	RLock()
	RUnlock()
}
