package main

import (
	"fmt"
	"github.com/hootuu/utils/configure"
)

func main() {
	fmt.Println(configure.GetString("sys.mode", "LOCAL01"))
}
