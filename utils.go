package gopool

import (
	"sync"

	"github.com/google/uuid"
	"github.com/golang-collections/collections/stack"
)

func readMap(mu *sync.Mutex, m map[Data]int, d Data) int {
	mu.Lock()
	defer mu.Unlock()

	return m[d]
}

func incrementMap(mu *sync.Mutex, m map[Data]int, d Data) {
	mu.Lock()
	defer mu.Unlock()

	m[d]++
}

func transform[V comparable](slice []V) (stck stack.Stack) {
	for i := len(slice) - 1; i >= 0; i-- {
		stck.Push(Data{uuid: uuid.New().String(), data: slice[i]})
	}

	return
}
