package controller

import (
	"back_end/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Weather struct {
}

func (w *Weather) Weather_get(c *gin.Context) {
	var CheckInfo struct {
		Lat string `json:"lat" form:"lat"`
		Lon string `json:"lon" form:"lon"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}

	err := srv.Weather.Weather_get(CheckInfo.Lat, CheckInfo.Lon)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}

	c.JSON(http.StatusOK, ResponseNew(c, "success!"))
}

func (w *Weather) Weather_broadcast(c *gin.Context) {
	err, resp := srv.Weather.Weather_broadcast()
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
