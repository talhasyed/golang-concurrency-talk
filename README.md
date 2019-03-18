theme: Next, 3

# Golang Concurrency

> Don't communicate by sharing memory; share memory by communicating

---

# Why is concurrency important

- multicore systems
- horizontal scaling
- tons of parallel workloads

---

# Traditional Concurrency Technicques

Threads, Mutexes, Semaphores, Promises, Futures, Callbacks, Oh My!

**Problems**: race conditions, deadlocks, and resource starvation, non parallelism, shared ownership of shared state

---

# 😵

---

# The golang way

> Don't communicate by sharing memory; share memory by communicating

But what does this mean?

---

# Our Program

**Please write a function to calculate the fibonnaci number ...🙀**

Ring any bells ;) ?

---

# A Sequential Implementation

[.code-highlight: none]

[.code-highlight: 1-7]

[.code-highlight: 9-19]

```golang
// Fibonacci calculates Fibonacci number
func Fib(n int) int {
	if n <= 1 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// Our main function
func calcFibsSequential() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		n := randInt(39, 41)
		result := Fib(n)

		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", i+1, n, result))
	}
}
```

---

## A Sequential Implementation

- Simple to write
- Easy to understand
- No concurrency bugs to worry about
- **SLOW!**

---

# Benchmarks

[.code-highlight: none]

[.code-highlight: all]

[.code-highlight: none]

```golang
func BenchmarkCalcFibsSequential(b *testing.B) {
	calcFibsSequential()
}
```

[.code-highlight: none]

[.code-highlight: all]

[.code-highlight: 28]

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/talhasyed/go-concurrency -bench ^(BenchmarkCalcFibsSequential)$

[1] 	fib(39) 	63245986
[2] 	fib(40) 	102334155
[3] 	fib(40) 	102334155
[4] 	fib(39) 	63245986
[5] 	fib(40) 	102334155
[6] 	fib(40) 	102334155
[7] 	fib(40) 	102334155
[8] 	fib(39) 	63245986
[9] 	fib(40) 	102334155
[10] 	fib(39) 	63245986
[11] 	fib(39) 	63245986
[12] 	fib(40) 	102334155
[13] 	fib(40) 	102334155
[14] 	fib(40) 	102334155
[15] 	fib(40) 	102334155
[16] 	fib(39) 	63245986
[17] 	fib(40) 	102334155
[18] 	fib(39) 	63245986
[19] 	fib(40) 	102334155
[20] 	fib(40) 	102334155
goos: darwin
goarch: amd64
pkg: github.com/talhasyed/go-concurrency
BenchmarkCalcFibsSequential-8   	       1	12287446420 ns/op	    2792 B/op	     107 allocs/op
PASS
ok  	github.com/talhasyed/go-concurrency	12.295s
Success: Benchmarks passed.
```

---

# Let's add concurrency

---

# go routines

- Lightweight green thread
- Managed by the go runtime
- Low latency of communication between goroutines
- Low cost to creation and teardown. Common to have millions!

---

# The `go` keyword

```golang
go Fib(n)
```

This runs `Fib()` in a separate go routine.

---

# Concurrent program

[.code-highlight: all]

[.code-highlight: 6]

[.code-highlight: 7]

[.code-highlight: all]

```golang
func calcFibsGoRoutine() {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 20; i++ {
		n := randInt(39, 41)
		go Fib(n)
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%s", i+1, n, "-"))
	}
}
```

---

# Results

[.code-highlight: all]

[.code-highlight: 28]

[.code-highlight: 7-25]

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/talhasyed/go-concurrency -bench ^(BenchmarkCalcFibsGoRoutine)$

goos: darwin
goarch: amd64
pkg: github.com/talhasyed/go-concurrency
BenchmarkCalcFibsGoRoutine-8   	[1] 	fib(39) 	-
[2] 	fib(40) 	-
[3] 	fib(40) 	-
[4] 	fib(40) 	-
[5] 	fib(39) 	-
[6] 	fib(40) 	-
[7] 	fib(39) 	-
[8] 	fib(40) 	-
[9] 	fib(40) 	-
[10] 	fib(40) 	-
[11] 	fib(40) 	-
[12] 	fib(40) 	-
[13] 	fib(39) 	-
[14] 	fib(39) 	-
[15] 	fib(40) 	-
[16] 	fib(40) 	-
[17] 	fib(39) 	-
[18] 	fib(39) 	-
[19] 	fib(39) 	-
[20] 	fib(40) 	-
       5	 224529616 ns/op	    1537 B/op	      20 allocs/op
PASS
ok  	github.com/talhasyed/go-concurrency	3.016s
Success: Benchmarks passed.
```

---

# 😕

---

**Problem**: How do we synchronize on the results of our go routines?

---

# go channels

A means to communicate over go routines.

By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

**This is what communicating sequential processes _is_ (CSP)**

---

# go channels

[.code-highlight: none]

[.code-highlight: 1-2]

[.code-highlight: 4-8]

[.code-highlight: 10-13]

```golang
// creating a channel
c := make(chan int)

// channel aware function
func calculate(n int, c chan int) {
	// value := ...
	c <- value
}

// invoking the channel aware function and
// receiving values from it
go calculate(n, c)
result := <- c
```

---

# Let's add go channels

[.code-highlight: all]

[.code-highlight: none]

[.code-highlight: 3]

[.code-highlight: 7,8]

[.code-highlight: all]

```golang
func calcFibChannels() {
	rand.Seed(time.Now().UnixNano())
	c := make(chan int, 20)

	for i := 0; i < 20; i++ {
		n := randInt(39, 41)
		go FibCalcChannel(n, c)
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", i+1, n, result))
	}
}
```

---

**Marvel** at the perf of _8 i7 cores_ doing **concurrent**, _**parallel**_ fibs!

👊🏼

---

[.code-highlight: 27]

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/talhasyed/go-concurrency -bench ^(BenchmarkCalcFibChannels)$

[1] 	fib(39) 	63245986
[2] 	fib(39) 	63245986
[3] 	fib(40) 	102334155
[4] 	fib(40) 	102334155
[5] 	fib(39) 	63245986
[6] 	fib(39) 	63245986
[7] 	fib(40) 	102334155
[8] 	fib(40) 	102334155
[9] 	fib(40) 	102334155
[10] 	fib(40) 	102334155
[11] 	fib(39) 	63245986
[12] 	fib(40) 	102334155
[13] 	fib(39) 	63245986
[14] 	fib(40) 	102334155
[15] 	fib(39) 	63245986
[16] 	fib(39) 	63245986
[17] 	fib(39) 	63245986
[18] 	fib(39) 	63245986
[19] 	fib(40) 	102334155
goos: darwin
goarch: amd64
pkg: github.com/talhasyed/go-concurrency
BenchmarkCalcFibChannels-8   	       1	10852114690 ns/op	    3768 B/op	     105 allocs/op
PASS
ok  	github.com/talhasyed/go-concurrency	10.860s
Success: Benchmarks passed.
```

---

# 😔

---

**Problem**: We need to concurrently receive the values of the fib channels.

---

Enough of the teasing; let's get it right this time!

---

## Concurrent go channel communication

[.code-highlight: none]

[.code-highlight: 1-6]

[.code-highlight: 12,17-22]

[.code-highlight: all]

```golang
// FibCalcJob respresents a job to calculate the Fib number of the value
type FibCalcJob struct {
	id     int
	value  int
	result int
}

func calcFibBlockingChannels() {
	rand.Seed(time.Now().UnixNano())
	NumCalcs := 20

	c := make(chan FibCalcJob, 20)

	for i := 0; i < NumCalcs; i++ {
		n := randInt(39, 41)
		go FibCalcChannel2(i+1, n, c)
	}

	for j := 0; j < NumCalcs; j++ {
		result := <-c
		fmt.Println(fmt.Sprintf("[%d] \tfib(%d) \t%d", result.id, result.value, result.result))
	}
	close(c)
}
```

---

[.code-highlight: none]

[.code-highlight: 3-22,28]

```
Running tool: /usr/local/bin/go test -benchmem -run=^$ github.com/talhasyed/go-concurrency -bench ^(BenchmarkCalcFibBlockingChannels)$

[4] 	fib(39) 	63245986
[2] 	fib(39) 	63245986
[19] 	fib(39) 	63245986
[7] 	fib(39) 	63245986
[12] 	fib(39) 	63245986
[15] 	fib(39) 	63245986
[3] 	fib(39) 	63245986
[13] 	fib(39) 	63245986
[14] 	fib(39) 	63245986
[17] 	fib(39) 	63245986
[20] 	fib(39) 	63245986
[10] 	fib(39) 	63245986
[6] 	fib(39) 	63245986
[1] 	fib(40) 	102334155
[16] 	fib(40) 	102334155
[5] 	fib(40) 	102334155
[18] 	fib(40) 	102334155
[9] 	fib(40) 	102334155
[8] 	fib(40) 	102334155
[11] 	fib(40) 	102334155
goos: darwin
goarch: amd64
pkg: github.com/talhasyed/go-concurrency
BenchmarkCalcFibBlockingChannels-8   	       1	2684620576 ns/op	   19336 B/op	     173 allocs/op
PASS
ok  	github.com/talhasyed/go-concurrency	2.692s
Success: Benchmarks passed.
```

---

# 🙌

---

> Don't communicate by sharing memory; share memory by communicating
> -- Rob Pike

- Don't have threads communicate by sharing a mermoy location with a data structure
- Have go routines (senders <-> receiver) share memory by message passing
- Use channels' blocking semantics (CSP) to ensure synchronization

---

Is there an even better way?

🤔
