package main

import (
	"fast_gin/core"
	"fmt"
)

func main() {
	cfg := core.ReadConfig()
	fmt.Printf("%v\n", cfg.DB)
}
