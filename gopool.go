package gopool

import (
	"sync"
	"errors"
)

func Pooler[V comparable](p PoolerParams[V]) {
	var (
		items    = transform(p.Slice)
		mu       = sync.Mutex{}
		errs     = make(map[Data]int)
		sem      = make(chan int, p.MaxRoutines)

		finished int
		done     bool
	)

	for !done {
		if items.Len() == 0 {
			continue
		}

		sem <- 1

		go func(item Data) {
			defer func() {
				<-sem
				done = len(p.Slice) <= finished
			}()

			var panicked bool

			if err := func() (e error) {
				defer func () {
					if r := recover(); r != nil {
						e = errors.New(r.(string))
						panicked = true
					}
				}()

				e = p.WorkerFn(item.data.(V))
				return

			}(); err != nil {
				incMap(&mu, errs, item)

				p.ErrorFn(err, panicked)

				if !(readMap(&mu, errs, item) >= p.MaxErrors && p.MaxErrors != -1) {
					items.Push(item)
				} else {
					finished++
				}

			} else {
				finished++
			}

		}(items.Pop().(Data))
	}
}
