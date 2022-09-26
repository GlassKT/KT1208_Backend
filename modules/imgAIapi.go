package modules

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func ImgAIapi(fileName string) bool {

	picture := map[string]string{
		"image": "http://10.80.161.247:8080/img/" + fileName, //*********** imgonly는 id + 확장자
	}

	imgJson, err := json.Marshal(picture)
	if err != nil {
		return false
	}

	buff := bytes.NewBuffer(imgJson)
	resp, err := http.Post("http://10.80.163.60:5000/image", "application/json", buff)
	if err != nil {
		return false
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false
	}

	respResult := string(respBody)

	if strings.Contains(respResult, "false") {
		return false
	}
	if strings.Contains(respResult, "error") {
		return false
	}

	return true
}
