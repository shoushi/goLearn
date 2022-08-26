package main

import (
	"github.com/gin-gonic/gin"
)

func mock(mockRout *gin.Engine) {
	group := mockRout.Group("/beetle/api")
	group.Use()
	{
		group.POST("inventoryBatchAdjustment/upload", func(ctx *gin.Context) {
			ctx.File("./wms反馈报文/批次调整反馈wms.json")
		})
	}
}
