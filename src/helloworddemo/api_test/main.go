package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main()  {

	for i:=0;i<1000;i++{
		go getHttp2()
	}

	time.Sleep(60*time.Second)
}

func getHttp2() {

	var url string = "https://admins.medtrib.cn/app/common/init/for_stat.json?sysDate=0&sysImei=0&userId=73256&sysSign=1d7af6601aaebeb304c71d7440417458"

	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "GolangSpider/1.0")
	//req.Header.Set("Host", "gltest.com") //incorrect way to set Host

	req.Host = "admins.medtrib.cn"

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("success")
	}

	log.Println(string(body))
}
