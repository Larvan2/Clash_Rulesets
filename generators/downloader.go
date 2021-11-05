package generators

import (
	"io/ioutil"
	"net/http"
)

func DownloadDirectList() {
	DirectUrl := "https://cdn.jsdelivr.net/gh/Loyalsoldier/v2ray-rules-dat@release/direct-list.txt"
	resp, err := http.Get(DirectUrl)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("direct.txt", data, 0644)
}
func DownloadBlockList() {
	CNBlockedUrl := "https://cdn.jsdelivr.net/gh/Loyalsoldier/v2ray-rules-dat@release/proxy-list.txt"
	resp, err := http.Get(CNBlockedUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("block.txt", data, 0644)
}

func DownloadAppleCNList() {
	apple_cnUrl := "https://cdn.jsdelivr.net/gh/Loyalsoldier/v2ray-rules-dat@release/apple-cn.txt"
	resp, err := http.Get(apple_cnUrl)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("apple-cn.txt", data, 0644)
}
