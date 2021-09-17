package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func Convertdirect() {
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
	directDomainlist, err := os.Open("direct.txt")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer directDomainlist.Close()

	os.Mkdir("output", 0755)
	output, _ := os.Create("output/direct.list")
	defer output.Close()

	inputReader := bufio.NewReader(directDomainlist)
	w := bufio.NewWriter(output)
	for {
		domain, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}
		if strings.Contains(string(domain), "full:") || strings.Contains(string(domain), "regexp:") {
			continue
		}
		w.WriteString("DOMAIN-SUFFIX," + string(domain) + "\n")
	}
}
