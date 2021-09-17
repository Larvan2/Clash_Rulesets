package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func Convertblocked() {
	CNBlockedUrl := "https://cdn.jsdelivr.net/gh/Loyalsoldier/cn-blocked-domain@release/domains.txt"

	//生成cnblock.list
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
	blockedDomainlist, err := os.Open("block.txt")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer blockedDomainlist.Close()

	os.Mkdir("output", 0755)
	output, _ := os.Create("output/cnblock.list")
	defer output.Close()

	inputReader := bufio.NewReader(blockedDomainlist)
	w := bufio.NewWriter(output)
	for {
		domain, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}

		w.WriteString("DOMAIN-SUFFIX," + string(domain) + "\n")
	}
}
