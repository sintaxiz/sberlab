package sbercloudAPI

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"example.com/cobraExample/sbercloudAPI/core"
)

var mainProjectId = "0ce45c2b0e8025e32f53c0067863b51b"

var key string
var secret string

func SetAuthData(k string, s string) {
	key = k
	secret = s
}

func signRequest(r *http.Request) {
	s := core.Signer{
		Key:    key,
		Secret: secret,
	}
	s.Sign(r)
}

func doGetRequestWithSigning(url string) *http.Response {

	request, err := http.NewRequest("GET", url, ioutil.NopCloser(bytes.NewBuffer([]byte(""))))
	request.Header.Add("X-Project-Id", "0ce45c2b0e8025e32f53c0067863b51b")
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	signRequest(request)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	return resp

}

func readBodyString(r *http.Response) string {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	return string(body)
}



func GetSubnets() string {
	url := "https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/" + mainProjectId + "/subnets"
	resp := doGetRequestWithSigning(url)
	return readBodyString(resp)
}

func GetEcs() string {
	url := "https://ecs.ru-moscow-1.hc.sbercloud.ru/v1/" + mainProjectId + "/cloudservers/detail?offset=1&limit=10"
	resp := doGetRequestWithSigning(url)
	return readBodyString(resp)

}

func GetVpcs() string {
	url := "https://vpc.ru-moscow-1.hc.sbercloud.ru/v1/" + mainProjectId + "/vpcs"
	resp := doGetRequestWithSigning(url)
	return readBodyString(resp)
}

func StartEcs(ecsId string) {
	url := "https://ecs.ru-moscow-1.hc.sbercloud.ru/v2.1/" + mainProjectId + "/servers/" + ecsId + "/action"
	request, err := http.NewRequest("POST", url, ioutil.NopCloser(bytes.NewBuffer([]byte("{\"os-start\": {}}"))))
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	request.Header.Add("os-stop", "")
	signRequest(request)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Println(string(body))
}

func StopEcs(ecsId string) {
	url := "https://ecs.ru-moscow-1.hc.sbercloud.ru/v2.1/" + mainProjectId + "/servers/" + ecsId + "/action"
	request, err := http.NewRequest("POST", url, ioutil.NopCloser(bytes.NewBuffer([]byte("{\"os-stop\": {}}"))))
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	request.Header.Add("os-stop", "")
	signRequest(request)

	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Printf("error: %s", err)
	}
	fmt.Println(readBodyString(resp))
}
