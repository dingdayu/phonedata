.PHONY: asset

asset:
	go-bindata -pkg phonedata -o bindata.go data

run:
	go run example/main.go

test:
	go test -v -bench=".*" -benchmem -memprofile memprofile.out -cpuprofile cpuprofile.out

tui:
	go tool pprof cpuprofile.out
