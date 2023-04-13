package gopool

import "errors"

func Pooler[V comparable](p PoolerParams[V]) {
	var (
		items    = transform(p.Slice)
		errs     = make(map[V]int)
		sem      = make(chan int, p.MaxRoutines)

		finished int
		done     bool
	)

	for !done {
		if items.Len() == 0 {
			continue
		}

		sem <- 1

		go func(item V) {
			defer func() {
				<-sem
				done = len(p.Slice) <= finished
			}()

			var panicked bool

			if err := func() (e error) {
				defer func (){
					if r := recover(); r != nil {
						e = errors.New(r.(string))
						panicked = true
					}
				}()

				e = p.WorkerFn(item)
				return

			}(); err != nil {
				errs[item]++

				p.ErrorFn(err, panicked)

				if !(errs[item] >= p.MaxErrors && p.MaxErrors != -1) {
					items.Push(item)
				} else {
					finished++
				}

			} else {
				finished++
			}

		}(items.Pop().(V))
	}
}
