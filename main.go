package main

import (
	"fmt"
	"runtime/debug"
)

func main() {
	fmt.Println(getVersion())
}

func getVersion() string {
	info, ok := debug.ReadBuildInfo()
	if !ok {
		return "(devel)"
	}
	return info.Main.Version
}
