package async

type NilType *struct{}

type ReturnType[T any] func() (T, error)

type AsyncFunc[Fn ReturnType[T], T any] func() Fn

// since golang uses interface implicitly, this line will provide static compile time check that ensure future always satisfied Awaitable
// the _ will tell the compiler to discard this line, so we can ensure there is no unneeded extra allocation
var _ Awaitable = &future[ReturnType[any], any]{}

type future[Fn ReturnType[T], T any] struct {
	sigend chan struct{}
	ret    Fn
}

// awaitSigend implements Awaitable
func (f *future[Fn, T]) awaitSigend() <-chan struct{} {
	return f.sigend
}

func (f *future[Fn, T]) Await() Fn {
	<-f.sigend
	return f.ret
}

func (f *future[Fn, T]) Retrieve() Fn {
	return f.ret
}

type runner[Fn ReturnType[T], T any] struct {
	future *future[Fn, T]
	runFn  AsyncFunc[Fn, T]
}

func (r *runner[Fn, T]) run() *future[Fn, T] {
	go func() {
		r.future.ret = r.runFn()
		r.future.sigend <- struct{}{}
	}()

	return r.future
}

func Go[Fn ReturnType[T], T any](f AsyncFunc[Fn, T]) *future[Fn, T] {
	r := &runner[Fn, T]{
		future: &future[Fn, T]{
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
