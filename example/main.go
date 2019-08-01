package main

import (
	"fmt"
	"github.com/dingdayu/phonedata"
)

func main() {
	info, _ := phonedata.Find("13298181006")
	fmt.Println(info)
}
