package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func downloadBlockList() {
	CNBlockedUrl := "https://cdn.jsdelivr.net/gh/Loyalsoldier/cn-blocked-domain@release/domains.txt"
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

	blockedDomainlist, err := os.Open("block.txt")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}

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

func convertBlockedTxt() {

	blockedDomainlist, err := os.Open("block.txt")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	os.Mkdir("output", 0755)
	output, _ := os.Create("output/cnblock.txt")
	defer output.Close()

	inputReader := bufio.NewReader(blockedDomainlist)
	w := bufio.NewWriter(output)
	for {
		domain, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}

		w.WriteString(string(domain) + "\n")
	}
}