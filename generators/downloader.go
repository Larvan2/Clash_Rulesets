package generators

import (
	"io/ioutil"
	"net/http"
)

func DownloadList(url, fileName string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("rules/"+fileName, data, 0644)
}
