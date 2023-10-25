package main

import (
	"github.com/hootuu/utils/configure"
	"github.com/hootuu/utils/sys"
)

func main() {
	mode := configure.GetString("sys.mode", "LOCAL01")
	sys.Info(sys.RunMode, " vs ", mode)
}
