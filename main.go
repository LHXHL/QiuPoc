package main

import (
	"PocSir/common"
	"time"
)

func main() {
	start := time.Now()
	common.FlagInfo()
	end := time.Since(start)
	common.TimePut.Printf("[*]运行时间为: %v\n", end)
}
