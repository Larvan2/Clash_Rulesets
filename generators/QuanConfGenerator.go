package generators

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ConvertDirectList_Quantumult() {

	directDomainlist, err := os.Open("rules/direct.txt")
	// apple_cn, err := os.Open("apple-cn.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer directDomainlist.Close()

	os.Mkdir("output", 0777)
	output, _ := os.Create("output/Quantumult.conf")
	defer output.Close()

	r := bufio.NewReader(directDomainlist)
	// ra := bufio.NewReader(apple_cn)
	w := bufio.NewWriter(output)

	w.WriteString("[SERVER]" + "\n\n")
	w.WriteString("[SOURCE]" + "\n" +
		"N1klaz Rules, filter, https://github.com/N1klaz/Clash_Rulesets/releases/latest/download/Quantumult.conf, true" + "\n\n")

	w.WriteString("[REWRITE] " + "\n" + "^https?://(www.)?(g|google).cn url 302 https://www.google.com" + "\n\n")
	w.WriteString("[TCP]" + "\n")

	w.WriteString("HOST-SUFFIX,local,DIRECT" + "\n" +
		"IP-CIDR,127.0.0.0/8,DIRECT" + "\n" +
		"IP-CIDR,172.16.0.0/12,DIRECT" + "\n" +
		"IP-CIDR,192.168.0.0/16,DIRECT" + "\n" +
		"IP-CIDR,10.0.0.0/8,DIRECT" + "\n" +
		"IP-CIDR,17.0.0.0/8,DIRECT" + "\n" +
		"IP-CIDR,100.64.0.0/10,DIRECT" + "\n" +
		"IP-CIDR,224.0.0.0/4,DIRECT" + "\n" + 
		"IP-CIDR6,fe80::/10,DIRECT" + "\n")

	w.WriteString("HOST-SUFFIX,cn,DIRECT" + "\n")

	for {
		if domain, _, err := r.ReadLine(); err == nil {
			if strings.Contains(string(domain), "regexp:") {
				continue
			} else if strings.Contains(string(domain), "full:") {
				w.WriteString(strings.Replace(string(domain), "full:", "HOST:", 1) + ",direct" + "\n")
			} else {
				w.WriteString("HOST-SUFFIX," + string(domain) + ",direct" + "\n")
			}
		} else {
			break
		}
	}

	w.WriteString("GEOIP,CN,DIRECT" + "\n" + "FINAL,PROXY" + "\n" + "[GLOBAL]" + "\n\n" + "[HOST]" + "\n\n" + "[STATE]" + "\n" +
		"STATE,AUTO" + "\n\n" + "[MITM]" + "\n")

	if err = w.Flush(); err != nil {
		fmt.Println(err)
	}

}
