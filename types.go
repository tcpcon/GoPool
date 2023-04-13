package gopool

type PoolerParams[V any] struct {
	Slice                  []V
	WorkerFn               func(V) error
	ErrorFn                func(error, bool)
	MaxRoutines, MaxErrors int
}
