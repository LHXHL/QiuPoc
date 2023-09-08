package common

import (
	"PocSir/config"
	"PocSir/scan"
	"flag"
	"github.com/fatih/color"
	"strings"
)

func Banner() {
	banner := `   ███████    ██         ███████                  
  ██░░░░░██  ░░         ░██░░░░██                 
 ██     ░░██  ██ ██   ██░██   ░██  ██████   █████ 
░██      ░██ ░██░██  ░██░███████  ██░░░░██ ██░░░██
░██    ██░██ ░██░██  ░██░██░░░░  ░██   ░██░██  ░░ 
░░██  ░░ ██  ░██░██  ░██░██      ░██   ░██░██   ██
 ░░███████ ██░██░░██████░██      ░░██████ ░░█████ 
  ░░░░░░░ ░░ ░░  ░░░░░░ ░░        ░░░░░░   ░░░░░
							by: Qiu
`
	color.HiMagenta(banner)
}
func Help() {
	help := `	
Usage: ./main -h 查看帮助信息
`
	color.HiRed(help)
}

func ShowPocs() {
	pocs := `
0: 用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件上传漏洞
1: 用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件读取漏洞
2: 亿赛通 电子文档安全管理系统 UploadFileFromClientServiceForClient 任意文件上传漏洞
3: 锐捷交换机 WEB 管理系统 EXCU_SHELL 信息泄露
4: 用友 畅捷通T+ GetStoreWarehouseByStore 远程命令执行漏洞
5: 任我行 CRM SmsDataList SQL注入漏洞
6: 企望制造 ERP comboxstore.action 远程命令执行漏洞
7: 金蝶OA 云星空 kdsvc 远程命令执行漏洞
8: 金和OA C6-GetSqlData.aspx SQL注入漏洞
9: JeecgBoot 企业级低代码平台 qurestSql SQL注入漏洞 CVE-2023-1454
10: 大华 智慧园区综合管理平台 video 任意文件上传漏洞
11: 大华 智慧园区综合管理平台 user_getUserInfoByUserName.action 账号密码泄漏漏洞
12: 红帆OA Ioffice Udfmr.asmx SQL注入漏洞
13: 蓝凌OA sysSearchMain.do 远程命令执行漏洞
...
`
	config.TextPut.Println(pocs)
}

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
		scan.PwnSingleTarget(*target, *expName)
	} else {
		Banner()
		scan.PwnTargets(*file, *expName, *thread)
	}
}
