package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertDirectList() {

	directDomainlist, err := os.Open("rules/direct.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer directDomainlist.Close()

	os.Mkdir("publish", 0777)
	output, _ := os.Create("publish/direct.list")
	defer output.Close()

	r := bufio.NewReader(directDomainlist)
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

func ConvertDirectTxt() {
	directDomainlist, err := os.Open("rules/direct.txt")

	if err != nil {
		fmt.Println(err)
	}
	output, _ := os.Create("publish/direct.txt")
	defer output.Close()

	r := bufio.NewReader(directDomainlist)
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
