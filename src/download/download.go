package download

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func Download(url string) (content string, err error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return "", err
	}
	
	resp, err := client.Do(req)
	defer resp.Body.Close()
	
	tcontent, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error", err.Error())
		return "", err
	}
	
	content = string(tcontent)
	return content, nil
}
