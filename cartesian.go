package cartesian

import (
	"sync"
)

// Iter takes interface-slices and returns a channel, receiving cartesian products
func Iter[K comparable, V any](params ...map[K][]V) chan map[K]V {
	// create channel
	c := make(chan map[K]V)
	// create waitgroup
	var wg sync.WaitGroup
	// call iterator
	wg.Add(1)

	// flatten any params passed in with multiple keys
	var combined []map[K][]V
	for _, m := range params {
		for key, val := range m {
			flattened := map[K][]V{
				key: val,
			}
			combined = append(combined, flattened)
		}
	}
	iterate(&wg, c, map[K]V{}, combined...)

	// call channel-closing go-func
	go func() { wg.Wait(); close(c) }()
	// return channel
	return c
}

// private, recursive Iteration-Function
func iterate[K comparable, V any](wg *sync.WaitGroup, channel chan map[K]V, result map[K]V, params ...map[K][]V) {
	// dec WaitGroup when finished
	defer wg.Done()
	// no more params left?
	if len(params) == 0 {
		// send result to channel
		channel <- result
		return
	}
	// shift first param
	p, params := params[0], params[1:]

	var pkey K
	for key := range p {
		pkey = key
		break
	}

	// iterate over it
	for i := 0; i < len(p[pkey]); i++ {
		// inc WaitGroup
		wg.Add(1)

		// create copy of result
		resultCopy := map[K]V{}
		for k, v := range result {
			resultCopy[k] = v
		}
		resultCopy[pkey] = p[pkey][i]

		// call self with remaining params
		go iterate[K, V](wg, channel, resultCopy, params...)
	}
}
