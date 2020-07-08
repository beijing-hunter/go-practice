package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"testing"
)

//go test -v -parallel 5 -count=1000
func TestHttp(t *testing.T) {
	t.Parallel() //并行运行测试用例，开头需要调用此函数

	var url string = "https://admins.medtrib.cn/app/user/get_info.json?sysDate=0&sysImei=0&userId=73256&sysSign=1d7af6601aaebeb304c71d7440417458"

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

func TestHttp2(t *testing.T) {

	t.Parallel() //并行运行测试用例，开头需要调用此函数
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

func TestHttp3(t *testing.T) {

	t.Parallel() //并行运行测试用例，开头需要调用此函数
	var url string = "https://admins.medtrib.cn/app/homepage/modules.json?listStatus=0&newConcernIds=18,104&pageOffset=0&pageSize=3&sysDate=0&sysImei=0&userId=73256&sysSign=ec808a3e5fcfd57ee87d188076eda5c0"

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

func TestHttp4(t *testing.T) {

	t.Parallel() //并行运行测试用例，开头需要调用此函数
	var url string = "https://admins.medtrib.cn/app/training/list.json?pageoffset=0&pagesize=20&selectCategoryId=-1&selectOrderById=1&selectPriceId=-1&sysDate=0&sysImei=0&userId=73256&sysSign=9ca44440d1ae4adcf19da702b29cfd3b"

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

func TestHttp5(t *testing.T) {

	t.Parallel() //并行运行测试用例，开头需要调用此函数
	var url string = "https://admins.medtrib.cn/app/training/list.json?pageoffset=0&pagesize=20&selectCategoryId=-1&selectOrderById=1&selectPriceId=-1&sysDate=0&sysImei=0&userId=73256&sysSign=9ca44440d1ae4adcf19da702b29cfd3b"

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

func TestHttp6(t *testing.T) {

	t.Parallel() //并行运行测试用例，开头需要调用此函数
	var url string = "https://admins.medtrib.cn/app/training/focus/tlist.json?focusUuid=1776c1ce-4664-4a33-af5b-aa902638eebc&sysDate=0&sysImei=0&userId=73256&sysSign=2cf8060bfe410580ba476733eb56d9f7"

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

func BenchmarkHttp(b *testing.B) {

	//b.Parallel() //并行运行测试用例，开头需要调用此函数
	resp, err := http.Get("https://admins.medtrib.cn/app/user/get_info.json?sysDate=0&sysImei=0&userId=73256&sysSign=1d7af6601aaebeb304c71d7440417458")

	if err == nil && resp.StatusCode == 200 {
		b.Log("success")
	} else {
		b.Error(err)
		b.Log("fail")
	}
}
