package main

import (
	"errors"
	"log"
	"time"

	"github.com/frojila/async"
)

func task(msg string) func() (int, error) {
	return func() (int, error) {
		time.Sleep(1 * time.Second)
		log.Println(msg)

		i := 15

		return i, errors.New("something bad happened")
	}
}

func main() {
	log.Println("start async task")
	future := async.Go(async.Wrap(task("hello world")).Get())

	log.Println("waiting task to finish")

	fn := future.Await()
	ret, err := fn()
	if err != nil {
		log.Println(err)
	}

	log.Println("task returning value: ", ret)

	log.Println("async task is finished")
}
