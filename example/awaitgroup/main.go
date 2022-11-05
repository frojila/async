package main

import (
	"log"
	"time"

	"github.com/frojila/async"
)

func task[Fn async.ReturnType[T], T int](msg string, number T) async.AsyncFunc[Fn, T] {
	return func() Fn {
		time.Sleep(10 * time.Millisecond)
		log.Println(msg)

		return func() (T, error) {
			return number, nil
		}
	}
}

func hello[Fn async.ReturnType[T], T async.NilType]() async.AsyncFunc[Fn, T] {
	return func() Fn {
		time.Sleep(1 * time.Second)
		log.Println("hello world")

		return nil
	}

}

func main() {
	log.Println("start async tasks")
	future1 := async.Go(task("hello from task one", 30))
	future2 := async.Go(hello())
	future3 := async.Go(hello())
	future4 := async.Go(task("hello from task two", 40))

	ag := async.AwaitGroup(future1, future2)
	ag.Add(future3)
	ag.Add(future4)

	log.Println("waiting all tasks to finish")

	ag.Await()

	fn1 := future1.Retrieve()
	ret1, _ := fn1()

	log.Println("task returning value: ", ret1)

	fn4 := future4.Retrieve()
	ret4, _ := fn4()

	log.Println("task returning value: ", ret4)

	log.Println("all async tasks is finished")
}
