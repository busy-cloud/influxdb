package apis

import (
	"github.com/busy-cloud/boat/api"
	"github.com/busy-cloud/boat/curd"
	"github.com/busy-cloud/boat/db"
	"github.com/busy-cloud/influxdb/influx2"
	"github.com/gin-gonic/gin"
)

func init() {
	api.Register("GET", "device/:id/history/:point", deviceHistory)
	api.Register("GET", "influxdb/device/:id/history/:point", deviceHistory)
}

type Device struct {
	Id        string `json:"id" xorm:"pk"`
	ProductId string `json:"product_id" xorm:"index"`
}

func deviceHistory(ctx *gin.Context) {
	var dev Device
	has, err := db.Engine.ID(ctx.Param("id")).Get(&dev)
	if err != nil {
		curd.Error(ctx, err)
		return
	}
	if !has {
		curd.Fail(ctx, "找不到设备")
		return
	}

	key := ctx.Param("point")
	start := ctx.DefaultQuery("start", "-5h")
	end := ctx.DefaultQuery("end", "0h")
	window := ctx.DefaultQuery("window", "10m")
	method := ctx.DefaultQuery("method", "mean") //last

	points, err := influx2.Query(dev.ProductId, dev.Id, key, start, end, window, method)
	if err != nil {
		curd.Error(ctx, err)
		return
	}

	curd.OK(ctx, points)
}
