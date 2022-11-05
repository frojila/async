package async

type NilType *struct{}

type ReturnType[T any] func() (T, error)

type AsyncFunc[Fn ReturnType[T], T any] func() Fn

type Future[Fn ReturnType[T], T any] struct {
	sigend chan struct{}
	ret    Fn
}

type runner[Fn ReturnType[T], T any] struct {
	future *Future[Fn, T]
	runFn  AsyncFunc[Fn, T]
}

func (r *Future[Fn, T]) Await() Fn {
	<-r.sigend
	return r.ret
}

func (r *runner[Fn, T]) run() *Future[Fn, T] {
	go func() {
		r.future.ret = r.runFn()
		r.future.sigend <- struct{}{}
	}()

	return r.future
}

func Go[Fn ReturnType[T], T any](f AsyncFunc[Fn, T]) *Future[Fn, T] {
	r := &runner[Fn, T]{
		future: &Future[Fn, T]{
			sigend: make(chan struct{}),
		},
		runFn: f,
	}

	r.run()

	return r.future
}

func Wrap[T any](fn ReturnType[T]) AsyncFunc[ReturnType[T], T] {
	return func() ReturnType[T] {
		return fn
	}
}
