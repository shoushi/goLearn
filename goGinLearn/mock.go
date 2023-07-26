package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func postMock(fileName string, group *gin.RouterGroup, dirPath string) {
	urlPath := strings.ReplaceAll(fileName, "_", "/")
	urlPath = strings.Split(urlPath, ".")[0]
	group.POST(urlPath, func(ctx *gin.Context) {

		// ctx.File(dirPath + fileName)
		ctx.Header("Content-Type", "application/json")
		ctx.String(200, generateMsg(dirPath+fileName))
	})
}
func mock(mockRout *gin.Engine, groupPath string, dirPath string) {
	group := mockRout.Group(groupPath)
	group.Use()
	{
		fileList, _ := os.ReadDir(dirPath)
		for _, file := range fileList {
			fmt.Println(file.Name())
			postMock(file.Name(), group, dirPath)
		}
	}
}

func generateMsg(fileName string) string {
	file, _ := os.OpenFile(fileName, os.O_RDWR, 0644)
	content, _ := ioutil.ReadAll(file)
	defer file.Close()

	currentTime := time.Now().Unix()
	timestampStr := strconv.FormatInt(currentTime, 10)
	newContent := strings.ReplaceAll(string(content), "${timestemp}", timestampStr)
	return newContent
}
