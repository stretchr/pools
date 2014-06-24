package pools

import "sync"

var once sync.Once

// Suite provides management of a suite of sync.Pool objects.
// Suite allows you to request a slice of objects of the type
// provided in the New function. The slice can be of any size
// and each size is backed by its own sync.Pool
type Suite struct {
	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}

	pools    map[int]*sync.Pool
	deadbolt sync.RWMutex
}

// Get selects an arbitrary item from the appropriate Pool, removes it from the Pool,
// and returns it to the caller. Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and the values returned by Get.
//
// If Get would otherwise return nil and p.New is non-nil, Get returns the result of calling p.New.
//
// The appropriate pool is chosen by the size provided.
func (s *Suite) Get(size int) []interface{} {
	once.Do(func() {
		if s.pools == nil {
			s.pools = map[int]*sync.Pool{}
		}
	})
	return s.getPool(size).Get().([]interface{})
}

// Put adds x to the appropriate pool.
//
// The appropriate pool is chosen by taking the length of the
// slice of interface{}.
func (s *Suite) Put(x []interface{}) {
	once.Do(func() {
		if s.pools == nil {
			s.pools = map[int]*sync.Pool{}
		}
	})
	s.getPool(int(len(x))).Put(x)
}

func (s *Suite) getPool(size int) *sync.Pool {

	s.deadbolt.RLock()
	pool, ok := s.pools[size]
	s.deadbolt.RUnlock()

	if !ok {
		pool = &sync.Pool{
			New: func() interface{} {
				return newSlice(size, s.New)
			},
		}
		s.deadbolt.Lock()
		s.pools[size] = pool
		s.deadbolt.Unlock()
	}
	return pool
}

func newSlice(size int, New func() interface{}) []interface{} {
	o := make([]interface{}, size)
	for i := 0; i < size; i++ {
		o[i] = New()
	}
	return o
}
