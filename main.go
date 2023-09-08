package main

import (
	"PocSir/common"
	"PocSir/config"
	"time"
)

func main() {
	start := time.Now()
	common.FlagInfo()
	end := time.Since(start)
	config.TimePut.Printf("[*]运行时间为: %v\n", end)
}
