package main

import (
	"fmt"
	"phonedata"
)

func main() {
	info, _ := phonedata.Find("13298181006")
	fmt.Println(info)
}
