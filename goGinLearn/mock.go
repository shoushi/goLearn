package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const dirPath string = "./wms反馈报文/"

func postMock(fileName string, group *gin.RouterGroup) {
	urlPath := strings.ReplaceAll(fileName, "_", "/")
	urlPath = strings.Split(urlPath, ".")[0]
	group.POST(urlPath, func(ctx *gin.Context) {
		ctx.File(dirPath + fileName)
	})
}
func mock(mockRout *gin.Engine) {
	group := mockRout.Group("/beetle/api")
	group.Use()
	{
		fileList, _ := os.ReadDir(dirPath)
		for _, file := range fileList {
			fmt.Println(file.Name())
			postMock(file.Name(), group)
		}
	}
}
