.PHONY: asset

bindata:
	go-bindata -ignore=data/bindata.go -pkg data -o data/bindata.go data

run:
	go run example/main.go

test:
	go test -v -bench=".*" -benchmem -memprofile memprofile.out -cpuprofile cpuprofile.out

tui:
	go tool pprof cpuprofile.out
