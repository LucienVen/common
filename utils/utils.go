package utils

import (
	"io/ioutil"
	"net/http"
)

// 获取公网IP
func GetPublicNetworkIp() string {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	content, _ := ioutil.ReadAll(resp.Body)
	return string(content)
}
