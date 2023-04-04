package main

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func generateJson(template string, resource string, count int, page int) string {
	json := ""
	for i := page * count; i < count*(page+1); i++ {
		temp := strings.ReplaceAll(template, "{num}", strconv.Itoa(i))
		if i == 0 {
			json = temp
		} else {
			json = json + "," + temp
		}
		println("now is " + strconv.Itoa(i))
	}
	resStr := strings.ReplaceAll(resource, "{template}", json)
	os.Create("./报文生成/result" + strconv.Itoa(page) + ".json")
	ioutil.WriteFile("./报文生成/result"+strconv.Itoa(page)+".json", []byte(resStr), 0664)
	return json
}

func main() {
	res1, _ := ioutil.ReadFile("./报文生成/resource.json")
	t1, _ := ioutil.ReadFile("./报文生成/template.json")
	template := string(t1)
	resource := string(res1)
	for i := 1; i < 99; i++ {
		go generateJson(template, resource, 100000, i)
	}
	generateJson(template, resource, 200000, 0)
}
