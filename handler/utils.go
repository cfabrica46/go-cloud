package handler

import (
	"io/ioutil"
)

func getAllFilesURL() (filesURL []string, err error) {
	files, err := ioutil.ReadDir("./cloud")
	if err != nil {
		return
	}

	for i := range files {
		filesURL = append(filesURL, "https://localhost:8081/cloud/"+files[i].Name())
	}
	return
}
