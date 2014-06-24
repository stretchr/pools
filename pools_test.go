package pools

import (
	"sync"

	"github.com/stretchr/testify/assert"

	"testing"
)

type obj struct {
}

func TestPools(t *testing.T) {
	var s Suite

	n := func() interface{} {
		return &obj{}
	}
	s.New = n

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 1000; j++ {
				x := s.Get(j)
				if assert.NotNil(t, s.pools) {
					s.deadbolt.RLock()
					if assert.NotNil(t, s.pools[j]) {
						assert.NotNil(t, s.pools[j].New)
					}
					s.deadbolt.RUnlock()
					assert.NotNil(t, x)
					assert.Equal(t, j, len(x))
				}
			}
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			for j := 0; j < 1000; j++ {
				x := s.Get(j)
				s.Put(x)
			}
		}
		wg.Done()
	}()

	wg.Wait()

}
