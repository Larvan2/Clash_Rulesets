package generators

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
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

func ConvertDirectList() {

	directDomainlist, err := os.Open("direct.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer directDomainlist.Close()

	os.Mkdir("output", 0777)
	output, _ := os.Create("output/direct.list")
	defer output.Close()

	r := bufio.NewReader(directDomainlist)
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

func ConvertDirectTxt() {
	directDomainlist, err := os.Open("direct.txt")

	if err != nil {
		fmt.Println(err)
	}
	output, _ := os.Create("output/direct.txt")
	defer output.Close()

	r := bufio.NewReader(directDomainlist)
	w := bufio.NewWriter(output)
	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "full:") || strings.Contains(string(domain), "regexp:") {
				continue
			} else {
				w.WriteString(string(domain) + "\n")
			}
		} else {
			break
		}
	}
	if err = w.Flush(); err != nil {
		fmt.Println(err)
	}
}
