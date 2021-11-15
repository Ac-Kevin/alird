package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Getnetip 获取本机外网IP
func Getnetip() (string, error) {
	urlStr := "https://api.ipify.org/?format=json"
	bytes, err := fetch(urlStr)
	if err != nil {
		return "", fmt.Errorf("获取ip失败: %s; url: %s", err, urlStr)
	}
	var result = map[string]string{}
	err = json.Unmarshal(bytes, &result)
	if err != nil {
		return "", fmt.Errorf("获取ip失败: %s; url: %s", err, urlStr)
	}
	ip := result["ip"]
	if ip != "" && ip != "127.0.0.1" {
		return ip, nil
	}
	return "", fmt.Errorf("获取ip失败: %s; url: %s", bytes, urlStr)
}

/*  fetch 网络请求，用于获取ip
传入参数说明: url 获取ip地址url
返回参数说明: byte数组，错误
*/
func fetch(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("http.fetch http.NewRequest error=%v\n", err))
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_12_0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Safari/537.36")
	var httpClient = http.Client{}
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("http.fetch httpClient.Do error=%v\n", err))
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(fmt.Sprintf("http.fetch wrong status code %d\n.", resp.StatusCode))
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf(fmt.Sprintf("读取http响应出错: %s\n", err))
	}
	return bytes, nil
}
