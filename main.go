package main

import "zcbs/cbs"

func main() {
	// download once a month is ok
	//all.DownloadAllCompanies()
	//coms := all.ParseCompaniesFromFileLine()
	//for _, c := range coms {
	//	fmt.Println(c)
	//}
	cbs.GetCBS("002597")


}
