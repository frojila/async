package async

type Awaitable interface {
	awaitSigend() <-chan struct{}
}

type awaitGroup struct {
	awaitables []Awaitable
}

func AwaitGroup(awaitables ...Awaitable) *awaitGroup {
	return &awaitGroup{
		awaitables: awaitables,
	}
}

func (w *awaitGroup) Add(awaitables Awaitable) {
	w.awaitables = append(w.awaitables, awaitables)
}

func (w *awaitGroup) Await() {
	for _, a := range w.awaitables {
		<-a.awaitSigend()
	}
}
