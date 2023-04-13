package gopool

import "github.com/golang-collections/collections/stack"

func transform[V comparable](slice []V) stack.Stack {
	var (
		new  []V
		stck stack.Stack
	)

	for _, v1 := range slice {
		if !func() bool {
			for _, v2 := range new {
				if v1 == v2 {
					return true
				}
			}

			return false
		}() {
			new = append(new, v1)
		}
	}

	for i := len(new) - 1; i >= 0; i-- {
		stck.Push(new[i])
	}

	return stck
}
