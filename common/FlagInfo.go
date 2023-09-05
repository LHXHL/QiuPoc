package common

import (
	"flag"
	"strings"
)

func FlagInfo() {
	target := flag.String("u", "", "目标URL")
	expName := flag.String("exp", "", "指定利用poc")
	file := flag.String("f", "", "指定测试的文件名")
	show := flag.Bool("show", false, "列出所有poc")
	thread := flag.Int("t", 30, "指定线程数")
	flag.Parse()

	*target = strings.TrimSuffix(*target, "/")

	if *target == "" && *file == "" && !*show {
		Banner()
		Help()
	} else if *show {
		Banner()
		ShowPocs()
	} else if *target != "" {
		Banner()
		PwnSingleTarget(*target, *expName)
	} else {
		Banner()
		PwnTargets(*file, *expName, *thread)
	}
}
