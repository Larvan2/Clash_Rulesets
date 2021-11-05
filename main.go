package main

import (
	"Clash_Rulesets/generators"
)

func main() {
	generators.DownloadBlockList()
	generators.DownloadDirectList()
	generators.DownloadAppleCNList()
	generators.ConvertBlockedList()
	generators.ConvertBlockedTxt()
	generators.ConvertDirectList()
	generators.ConvertDirectTxt()
	generators.ConvertDirectList_Quantumult()
}
