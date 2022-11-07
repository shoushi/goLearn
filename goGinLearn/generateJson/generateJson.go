package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func generateJson(count int) string {
	res, _ := ioutil.ReadFile("./报文生成/resource.json")
	t1, _ := ioutil.ReadFile("./报文生成/template.json")
	json := ""
	template := string(t1)
	for i := 0; i < count; i++ {
		temp := strings.ReplaceAll(template, "{num}", strconv.Itoa(i))
		if i == 0 {
			json = temp
		} else {
			json = json + "," + temp
		}
	}
	resStr := strings.ReplaceAll(string(res), "{template}", json)
	ioutil.WriteFile("./报文生成/result.json", []byte(resStr), 0664)
	return resStr
}

func main() {
	generateJson(2000)
}
