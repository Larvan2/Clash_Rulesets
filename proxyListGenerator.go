package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func downloadBlockList() {
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
func convertBlockedList() {
	proxyDomainlist, err := os.Open("block.txt")
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("output", 0755)
	output, _ := os.Create("output/proxy.list")

	defer output.Close()

	r := bufio.NewReader(proxyDomainlist)
	w := bufio.NewWriter(output)

	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "full:") || strings.Contains(string(domain), "regexp:") {
				continue
			} else {
				w.WriteString("DOMAIN-SUFFIX," + string(domain) + "\n")
			}
		} else {
			break
		}
	}
	if err = w.Flush(); err != nil {
		fmt.Println(err)
	}

}

func convertBlockedTxt() {
	proxyDomainList, err := os.Open("block.txt")
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("output", 0755)
	output, _ := os.Create("output/proxy.txt")

	defer output.Close()

	r := bufio.NewReader(proxyDomainList)
	w := bufio.NewWriter(output)

	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "full:") || strings.Contains(string(domain), "regexp:") {
				continue
			} else {
				w.WriteString("DOMAIN-SUFFIX," + string(domain) + "\n")
			}
		} else {
			break
		}
	}
	if err = w.Flush(); err != nil {
		fmt.Println(err)
	}
}
