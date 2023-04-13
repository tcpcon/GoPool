package gopool

import "github.com/golang-collections/collections/stack"

func transform[V any](slice []V) (stck stack.Stack) {
	for i := len(slice) - 1; i >= 0; i-- {
		stck.Push(slice[i])
	}

	return
}
