package controller

import (
	"back_end/common"
	"back_end/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Schedule struct {
}

func (s *Schedule) Class_add(c *gin.Context) {
	CheckInfo := model.Class{}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.OpErr))
		return
	}
	err := srv.Schedule.Class_add(CheckInfo)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, err))
}

func (s *Schedule) Class_broadcast(c *gin.Context) {
	var CheckInfo struct {
		Day int `json:"day" form:"day"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.OpErr))
		return
	}
	err, resp := srv.Schedule.Class_broadcast(CheckInfo.Day)
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}

func (s *Schedule) Class_get(c *gin.Context) {
	var CheckInfo struct {
		Day   int `json:"day" form:"day"`
		Class int `json:"time" form:"time"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.OpErr))
		return
	}
	resp := srv.Schedule.Class_get(CheckInfo.Day, CheckInfo.Class)
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
