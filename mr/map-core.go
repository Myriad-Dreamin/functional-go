package mr

import "sync"

type Handler interface {
	Call(f interface{}, index int)
}

func mapSlice(handler Handler, l, r int, f interface{}) {
	for i := l; i < r; i++ {
		handler.Call(f, i)
	}
}

func min(l, r int) int {
	if l < r {
		return l
	}
	return r
}

func MapSlice(handler Handler, l, r, coreCount int, f interface{}) Handler {
	step := (r - l + coreCount - 1) / coreCount
	if coreCount <= 1 || step <= 0 {
		mapSlice(handler, l, r, f)
	} else {
		var wg sync.WaitGroup
		wg.Add(coreCount)
		for i := l; i < r; i += step {
			go func(i int) {
				mapSlice(handler, i, min(i+step, r), f)
				wg.Done()
			}(i)
		}
		wg.Wait()
	}
	return handler
}
