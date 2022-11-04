# Async

Async is goroutine wrapper for golang with async/await manner

## Requirements
```
Go 1.18
```

## How to install
```bash
go get github.com/frojila/async
```

## Example usage
```go
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
```
## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

Please make sure to update tests as appropriate.

## License
[MIT](https://choosealicense.com/licenses/mit/)