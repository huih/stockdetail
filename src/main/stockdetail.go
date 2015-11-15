package main

import (
	"fmt"
	"encoding/json"
	"strings"
	"download"
	"url"
	"time"
)

type ipeople interface{
	GetName() string
};

type iperson struct {
	name string
};

func (p iperson) GetName() string{
	return p.name
}

func main(){
	code := "600812"
	
	type pageData struct{
			Pages int
			Data []string
		}
		
	var pagedata pageData
	var curpage int;
	var lastdata string
	
	curpage = 1
	for j := 0; j >= 0; j++ {
			
		for i := 0; i >= 0; i++ {
			url, err := url.Url(curpage, code, "")
			if err != nil {
				fmt.Println("url error", err.Error())
				return;
			}
		
			//fmt.Println("url:", url)
			content, err := download.Download(url)
			if err != nil {
				fmt.Println("download data error", err.Error())
				return;
			}
			
			datastr := string(content)
			
			title := "var jsTimeSharingData="
			datastr = datastr[len(title): len(datastr)]
			datastr = strings.Replace(datastr, "pages", "\"pages\"", -1)
			datastr = strings.Replace(datastr, ",data", ",\"data\"", -1)
			datastr = strings.Replace(datastr, ";", "", -1)
		
			dataBytes := []byte(datastr)
			err = json.Unmarshal(dataBytes, &pagedata)
			if err != nil {
				fmt.Println("Fatal error", err.Error())
			}
			
			if pagedata.Pages <= 1 || (curpage >= pagedata.Pages) {
				break;
			}
			curpage = pagedata.Pages;
		}
		
		if len(pagedata.Data) <= 0  {
			fmt.Println("current no data")
		} else {
			count := 0
			for i := len(pagedata.Data) - 10; i < len(pagedata.Data) && count <= 10; i++ {
				tmpdata := strings.Split(pagedata.Data[i], ",")
				if strings.Compare(tmpdata[0], lastdata) <= 0 {
					continue
				}
				if len(tmpdata) >= 4 && tmpdata[3] == "1" {
					fmt.Println(tmpdata[0], tmpdata[1], tmpdata[2], "买")
				} else {
					fmt.Println(tmpdata[0], tmpdata[1], tmpdata[2], "卖")
				}
				count = count +  1;
				lastdata = tmpdata[0]
			}
		} 	
		time.Sleep(5000 * time.Millisecond)
	}
}
		
	

