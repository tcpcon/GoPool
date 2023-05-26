# GoPool
Simplistic goroutine pooling package to complete a set of tasks within a certain amount of concurrent goroutines + handle errors.

```go
gopool.Pooler(gopool.PoolerParams[int]{Slice: []int{1, 2, 3, 4, 5, 6}, WorkerFn: func(item int) error {
	time.Sleep(1 * time.Second)
	println(item)
	  
	return nil
		  
}, ErrorFn: func(item int, err error, isPanic bool) {
	if isPanic {
		log.Error("Caught Panic: %s", err).Full()
	} else {
		log.Error(err.Error()).Msg()
	}
}, MaxRoutines: 2, MaxErrors: 1})
```

## Params
- `Slice`, items to complete the task upon, will act on them in an orderly fashion
- `WorkerFn`, a functon taking an item of the same type (**comparable**) as the list elements that will perform the task on each item, can return an error which will be handled by the below function
- `ErrorFn`, a function that must **not** panic that should log any errors or panics that are caught from the `WorkerFn`
- `MaxRoutines`, is the max concurrent goroutines for the worker pool
- `MaxErrors`, is the maximum amount of errors for an item in a `WorkerFn`
