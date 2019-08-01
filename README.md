# phonedata

> note: data from network, please use after identification.

<p style="text-align: center">
    <a href="https://travis-ci.org/dingdayu/phonedata"><img src="https://travis-ci.org/dingdayu/phonedata.svg?branch=master" alt="Build Status"></a>
    <a href="https://codeclimate.com/github/dingdayu/phonedata/maintainability"><img src="https://api.codeclimate.com/v1/badges/b9a238d60b2cb676d8ea/maintainability" alt="Code maintainability" /></a>
    <a href="https://codeclimate.com/github/dingdayu/phonedata/test_coverage"><img src="https://api.codeclimate.com/v1/badges/b9a238d60b2cb676d8ea/test_coverage" alt="code coverage"/></a>
    <a href="https://ci.appveyor.com/project/dingdayu/phonedata"><img src="https://ci.appveyor.com/api/projects/status/github/dingdayu/phonedata?svg=true&branch=master&passingText=Windows%20-%20OK&failingText=Windows%20-%20failed&pendingText=Windows%20-%20pending" alt="Windows Build Status"></a>
    <a href="https://blog.dingxiaoyu.com"><img src="https://img.shields.io/badge/author-@dingdayu-blue.svg?style=flat" alt="Author"></a>
    <a href="https://godoc.org/github.com/dingdayu/phonedata"><img src="https://godoc.org/github.com/dingdayu/phonedata?status.svg" alt="GoDoc"></a>
    <a href="https://goreportcard.com/report/github.com/dingdayu/phonedata"><img src="https://goreportcard.com/badge/github.com/dingdayu/phonedata" alt="Report"></a>
</p>

## install

```bash
go get -u github.com/dingdayu/phonedata
```

## example

```go
package main

import (
	"fmt"
	"phonedata"
)

func main() {
	info, _ := phonedata.Find("13298181006")
	fmt.Println(info)
}
```

> at go 1.12, go models

file: [example/main.go](example/main.go)

## benchmark

```shell
go test -v -bench=".*" -benchmem -memprofile memprofile.out -cpuprofile cpuprofile.out
=== RUN   TestIsPhone
--- PASS: TestIsPhone (0.00s)
=== RUN   TestTotalRecord
--- PASS: TestTotalRecord (0.00s)
=== RUN   TestVersion
--- PASS: TestVersion (0.00s)
=== RUN   TestLoadDataFile
--- PASS: TestLoadDataFile (0.00s)
=== RUN   Example
--- PASS: Example (0.00s)
goos: darwin
goarch: amd64
pkg: phonedata
BenchmarkIsPhone-4        200000             13505 ns/op            9209 B/op        114 allocs/op
BenchmarkVersion-4      300000000                5.21 ns/op            0 B/op          0 allocs/op
BenchmarkFind-4          3000000               449 ns/op             224 B/op          6 allocs/op
PASS
ok      phonedata       6.922s
```

at: MacBook Pro(Intel Core i7/3.1 GHz)

## Reference

The project is affected by the following projects or articles.

- golang [https://github.com/xluohome/phonedata](https://github.com/xluohome/phonedata)
- python: [https://github.com/lovedboy/phone](https://github.com/lovedboy/phone)
- php: [https://github.com/shitoudev/phone-location](https://github.com/shitoudev/phone-location)