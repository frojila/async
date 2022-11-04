package main

import (
	"log"
	"time"

	"github.com/frojila/async"
)

func task[Fn async.ReturnType[T], T int](msg string) async.AsyncFunc[Fn, T] {
	return func() Fn {
		time.Sleep(10 * time.Millisecond)
		log.Println(msg)

		var i T = 10

		return func() (T, error) {
			return i, nil
		}
	}
}

func main() {
	log.Println("start async task")
	future := async.Go(task("hello world"))

	log.Println("waiting task to finish")

	fn := future.Await()
	ret, _ := fn()

	log.Println("task returning value: ", ret)

	log.Println("async task is finished")
}
