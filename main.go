package main

import (
	"PocSir/config"
	"PocSir/poc"
	"bufio"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
	"sync"
	"time"
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
...
`
	config.TextPut.Println(pocs)
}

type Pocs func(string)

var PocsMap = map[string]Pocs{
	"0":  poc.U8CRM_upload_exp,
	"1":  poc.U8CRM_pathTravel_exp,
	"2":  poc.YiSaiTong_upload_Exp,
	"3":  poc.Ruijie_Excu_Shell,
	"4":  poc.Tplus_RCE,
	"5":  poc.RenWoXin_Crm_Sql,
	"6":  poc.QiWangERP_EXEC,
	"7":  poc.Kingdee_Erp_Kdsvc_RCE,
	"8":  poc.JinHeSql_Exec,
	"9":  poc.JeecgBoot_Sql,
	"10": poc.DaHua_Video_Upload,
	"11": poc.DaHua_sys_user,
	"99": func(target string) {
		fmt.Println("exit")
	},
}

func PwnSingleTarget(target string, expName string) {
	pwnFunc, found := PocsMap[expName]
	if !found {
		config.ErrMsg.Printf("[-]没有此序列的poc")
		return
	}
	pwnFunc(target)
}

func PwnTargets(file string, expName string, thread int) {
	urls, err := os.Open(file)
	if err != nil {
		config.ErrMsg.Println("[-]打开urls文件失败")
		return
	}
	defer urls.Close()
	var wg sync.WaitGroup
	urlChan := make(chan string)

	for i := 0; i < thread; i++ {
		wg.Add(1)
		go goUrl(urlChan, &wg, expName)
	}

	scanner := bufio.NewScanner(urls)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			urlChan <- line
		}
	}
	close(urlChan)
	wg.Wait()
}

func goUrl(urlChan <-chan string, wg *sync.WaitGroup, expName string) {
	defer wg.Done()
	for target := range urlChan {
		PwnSingleTarget(target, expName)
	}

}

func main() {
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
		start := time.Now()
		PwnTargets(*file, *expName, *thread)
		end := time.Since(start)
		config.TimePut.Printf("[*]运行时间为: %v\n", end)
	}
}
