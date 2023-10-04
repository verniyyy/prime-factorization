package primenum

import (
	"sync"
	"sync/atomic"
)

var (
	max          = atomic.Int64{}
	primeNumbers = sync.Map{}
	v            = struct{}{}
)

func GetMaxInCache() int64 {
	return max.Load()
}

// CacheClear ...
func CacheClear() {
	primeNumbers.Range(func(k, _ any) bool {
		primeNumbers.Delete(k)
		return true
	})
}

// CacheStore ...
func CacheStore(s ...int64) {
	for _, n := range s {
		primeNumbers.Store(n, v)
	}
	if s[len(s)-1] > max.Load() {
		max.Store(s[len(s)-1])
	}
}

// CacheExists ...
func CacheExists(n int64) bool {
	_, ok := primeNumbers.Load(n)
	return ok
}

// CacheEach ...
// If f returns false, range stops the iteration.
func CacheEach(f func(n int64) bool) {
	primeNumbers.Range(func(key, value any) bool {
		return f(key.(int64))
	})
}
