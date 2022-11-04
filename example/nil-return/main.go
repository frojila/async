package main

import (
	"log"
	"time"

	"github.com/frojila/async"
)

func task[Fn async.ReturnType[T], T async.NilType]() async.AsyncFunc[Fn, T] {
	return func() Fn {
		time.Sleep(1 * time.Second)
		log.Println("hello world")

		return nil
	}

}

func main() {
	log.Println("start async task")
	future := async.Go(task())

	log.Println("waiting task to finish")

	future.Await()

	log.Println("async task is finished")
}
