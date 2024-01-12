package tools

import (
	"runtime"
	"sync"
)

func innerMap[T any](data []T, f func(T) T, bidx, eidx int, wg *sync.WaitGroup) {
	defer wg.Done()
	if eidx == -1 {
		eidx = len(data)
	}
	for i := bidx; i < eidx; i++ {
		data[i] = f(data[i])
	}
}

func Map[T any](data []T, f func(T) T, num_worker ...int) {
	length := len(data)
	if length == 0 {
		return
	}
	var n int
	if len(num_worker) == 1 {
		n = num_worker[0]
	} else {
		n = runtime.NumCPU()
	}
	if length < n {
		n = length
	}

	inc := length / n
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		bidx := inc * i
		eidx := bidx + inc
		if i == n - 1 {
			wg.Add(1)
			go innerMap(data, f, bidx, -1, &wg)
		} else {
			wg.Add(1)
			go innerMap(data, f, bidx, eidx, &wg)
		}
	}
	wg.Wait()
}
