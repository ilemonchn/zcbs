package all

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
	"zcbs/comm"
)

type Company struct {
	Name         string
	Code         string
	CurrentPrice float32
}

type Resp struct {
	Data Data `json:"data,omitempty"`
}

type Data struct {
	Total int `json:"total,omitempty"`
	Diff []Coms `json:"diff,omitempty"`
}

type Coms struct {
	F2 float32 `json:"f2,omitempty"`
	F12 string `json:"f12,omitempty"`
	F14 string `json:"f14,omitempty"`
}

func DownloadAllCompanies() []Company {
	pageNo := 1
	pageSize := 100
	total := 4304
	var allComs []Company
	// first page
	total, coms := getPageComs(pageNo, pageSize)
	pageNo++
	allComs = append(allComs, coms...)
	for pageNo <= total/pageSize+1 {
		time.Sleep(2*time.Second)
		fmt.Println("pageNo", pageNo)
		_, coms = getPageComs(pageNo, pageSize)
		allComs = append(allComs, coms...)
		pageNo++
	}
	fmt.Println("total", total)
	buf := &bytes.Buffer{}
	for _, c := range allComs {
		buf.WriteString(fmt.Sprintf("%s, %s, %.2f\n", c.Code, c.Name, c.CurrentPrice))
	}
	file, err := os.OpenFile(CompaniesFilePath, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		n, err := file.Write(buf.Bytes())
		fmt.Println(n, err)
	}
	return nil
}

func getPageComs(pageNo, pageSize int) (int, []Company) {
	path := CompaniesPath
	path = fmt.Sprintf(path, pageNo, pageSize)
	// f2 12 14
	response, err := comm.HttpGet(path)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		fmt.Println(response.Status)
	}
	b, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}
	r := Resp{}
	// some price is - will get unmarshal string to float err
	err = json.Unmarshal(b, &r)
	if err != nil {
		fmt.Println(err)
	}
	var coms []Company
	for _, c := range r.Data.Diff {
		com := Company{
			Name: c.F14,
			Code: c.F12,
			CurrentPrice: c.F2,
		}
		coms = append(coms, com)
	}
	return r.Data.Total, coms
}

func ParseCompaniesFromFileLine() []Company {
	file, err := os.OpenFile(CompaniesFilePath, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	lines := strings.Split(string(all), "\n")
	cs := make([]Company, 0)
	for _, l := range lines {
		fs := strings.Split(l, ",")
		if len(fs) != 3 {
			continue
		}
		c := Company{}
		c.Code = strings.Trim(fs[0], " ")
		c.Name = strings.Trim(fs[1], " ")
		price, _ := strconv.ParseFloat(strings.Trim(fs[2], " "), 32)
		c.CurrentPrice = float32(price)
		cs = append(cs, c)
	}
	return cs
}
