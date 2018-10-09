### goroutine pool


Testing

```
=== RUN   TestInit
--- PASS: TestInit (0.00s)
=== RUN   TestGo
--- PASS: TestGo (1.00s)
    pool_test.go:46: 1 1
=== RUN   TestNew
--- PASS: TestNew (0.00s)
PASS
coverage: 89.5% of statements

```
Benchmark

```
goos: darwin
goarch: amd64
pkg: gpool
BenchmarkGo-4    	 5000000	       235 ns/op
BenchmarkGo2-4   	  300000	      4406 ns/op  [with io]
PASS
coverage: 84.2% of statements

Process finished with exit code 0

```


### Install

```
go get -u github.com/clearcodecn/gpool
```

### Usage

#### for default settings: `minSize`:`100` ,`maxSize`: `1000`

```
    g := gpool.New(1000,2000)

    // run task
    g.Go(func(){
        // your code .
    })

    // stop the gpool

    g.Stop()

    // ------------
    // use the default settings.
    gpool.Init()

    gpool.Run(func(){
        // your code
    })

    gpool.Stop()

```







