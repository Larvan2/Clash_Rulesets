package generators

import (
	"io/ioutil"
	"net/http"
	"os"
)

func DownloadList(url, fileName string) {
	os.Mkdir("rules", 0777)
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
