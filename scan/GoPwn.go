package scan

import (
	"PocSir/config"
	"PocSir/poc"
	"bufio"
	"fmt"
	"os"
	"sync"
)

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
	"12": poc.HongFan_Ioffice_Sql,
	"13": poc.LandrayOA_Rce,
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
