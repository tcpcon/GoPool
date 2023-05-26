package gopool

type (
	PoolerParams[V comparable] struct {
		Slice                  []V
		WorkerFn               func(V) error
		ErrorFn                func(V, error, bool)
		MaxRoutines, MaxErrors int
	}

	Data struct {
		uuid string
		data any
	}
)
