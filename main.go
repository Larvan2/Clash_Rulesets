package main

import (
	"Clash_Rulesets/generators"
)

func main() {
	generators.DownloadBlockList()
	generators.DownloadDirectList()
	generators.ConvertBlockedList()
	generators.ConvertBlockedTxt()
	generators.ConvertDirectList()
	generators.ConvertDirectTxt()
	// downloadBlockList()
	// downloadDirectList()
	// convertBlockedList()
	// convertDirectList()
	// convertBlockedTxt()
	// convertDirectTxt()
}
