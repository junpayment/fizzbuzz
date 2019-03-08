# FizzBuzz
You can write fizzbuzz quickly.

# Installation

```sh
go get github.com/junpayment/fizzbuzz
```

# DoSync/DoAsync
You can write fizzbuzz sync or async modes.

## sync

```go
package main

import (
	"fmt"
	
	"github.com/junpayment/fizzbuzz"
)

func main() {
	sets := fizzbuzz.DoSync(100)
	for _, set := range sets {
		fmt.Printf("%d : %s\n", set.Input, set.Res)
	}
}
```

## async

```go
package main

import (
	"fmt"
	
	"github.com/junpayment/fizzbuzz"
)

func main() {
	sets := fizzbuzz.DoAsync(100, 10)
	for _, set := range sets {
		fmt.Printf("%d : %s\n", set.Input, set.Res)
	}
}
```

enjoy!

# Benchmark


Since fizzbuzz is processing with minimal overhead, the parallel processing context seems to take longer than usual.

I recommend `DoSync`.

```
goos: darwin
goarch: amd64
pkg: fizzbuzz
BenchmarkDo/num_10000/do_sync-8 	    2000	    707744 ns/op
BenchmarkDo/num_10000/do_async/goroutine_1-8         	     200	   8040204 ns/op
BenchmarkDo/num_10000/do_async/goroutine_5-8         	     200	   6065541 ns/op
BenchmarkDo/num_10000/do_async/goroutine_10-8        	     200	   5985554 ns/op
BenchmarkDo/num_10000/do_async/goroutine_20-8        	     300	   5787948 ns/op
BenchmarkDo/num_10000/do_async/goroutine_100-8       	     300	   5782902 ns/op
BenchmarkDo/num_10000/do_async/goroutine_1000-8      	     200	   6126495 ns/op
```


