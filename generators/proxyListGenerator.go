package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertBlockedList() {
	proxyDomainlist, err := os.Open("rules/proxy.txt")
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("publish", 0755)
	output, _ := os.Create("publish/proxy.list")

	defer output.Close()

	r := bufio.NewReader(proxyDomainlist)
	w := bufio.NewWriter(output)

	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "regexp:") {
				continue
			} else if strings.Contains(string(domain), "full:") {
				w.WriteString(strings.Replace(string(domain), "full:", "DOMAIN,", 1) + "\n")
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

func ConvertBlockedTxt() {
	proxyDomainList, err := os.Open("rules/proxy.txt")
	if err != nil {
		fmt.Println(err)
	}
	os.Mkdir("output", 0755)
	output, _ := os.Create("publish/proxy.txt")

	defer output.Close()

	r := bufio.NewReader(proxyDomainList)
	w := bufio.NewWriter(output)

	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "regexp:") {
				continue
			} else if strings.Contains(string(domain), "full:") {
				w.WriteString(strings.Replace(string(domain), "full:", "", 1) + "\n")
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
