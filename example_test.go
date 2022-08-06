package mergeflight_test

import (
	"fmt"
	"sync"
	"time"

	"github.com/kkHAIKE/mergeflight"
)

func Example() {
	batchFunc := func(args []interface{}) (interface{}, error) {
		// simple multiply by 2, and returns map
		ret := make(map[int]int)
		for _, v := range args {
			ret[v.(int)] = v.(int) * 2
		}
		return ret, nil
	}

	// set count window to 5, time window to 10ms
	m := mergeflight.New(5, 10*time.Millisecond)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			ret, _ := m.Do(i, batchFunc)
			fmt.Println(ret)
		}(i)
	}
	wg.Wait()

	// Output: map[0:0 1:2 2:4 3:6 4:8]
	// map[0:0 1:2 2:4 3:6 4:8]
	// map[0:0 1:2 2:4 3:6 4:8]
	// map[0:0 1:2 2:4 3:6 4:8]
	// map[0:0 1:2 2:4 3:6 4:8]
}
