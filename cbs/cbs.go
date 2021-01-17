package cbs

import (
	"fmt"
	"io/ioutil"
	"os"
	"zcbs/comm"
)

const (
	cbsPath = "https://caibaoshuo.com/companies/%s"
)

type CBS struct {
	code string
	cbsLow int
	cbsHigh int
	ROE float32
	year int
}

func GetCBS(code string)  {
	path := fmt.Sprintf(cbsPath, code)
	resp, err := comm.HttpGet(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	filePath := fmt.Sprintf("./cbs/allhtml/%s.html", code)

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		n, err := file.Write(all)
		fmt.Println(n, err)
	}
}
