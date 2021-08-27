package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	Url := "https://cdn.jsdelivr.net/gh/Loyalsoldier/cn-blocked-domain@release/domains.txt"
	resp, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	ioutil.WriteFile("domains.txt", data, 0644)

	domainlist, err := os.Open("domains.txt")
	if err != nil {
		fmt.Println(err)
		return // exit the function on error
	}
	defer domainlist.Close()

	os.Mkdir("output", 0755)
	output, _ := os.Create("output/cnblock.list")
	defer output.Close()

	inputReader := bufio.NewReader(domainlist)
	w := bufio.NewWriter(output)
	for {
		domain, _, err := inputReader.ReadLine()
		if err == io.EOF {
			return
		}

		w.WriteString("DOMAIN-SUFFIX," + string(domain) + "\n")
	}

}
