package main

import (
	"Clash_Rulesets/generators"
)

func main() {
	direct := "https://cdn.jsdelivr.net/gh/Loyalsoldier/v2ray-rules-dat@release/direct-list.txt"
	proxy := "https://cdn.jsdelivr.net/gh/Loyalsoldier/v2ray-rules-dat@release/proxy-list.txt"
	generators.DownloadList(direct, "direct.txt")
	generators.DownloadList(proxy, "proxy.txt")
	generators.ConvertBlockedList()
	generators.ConvertBlockedTxt()
	generators.ConvertDirectList()
	generators.ConvertDirectTxt()
	generators.ConvertDirectList_Quantumult()
}
