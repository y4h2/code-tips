package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestConcurrency(t *testing.T) {
	var wg sync.WaitGroup
	n := 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			fmt.Printf("%p\n", GetLoadBalancer())
			wg.Done()
		}()
	}

	wg.Wait()
}
