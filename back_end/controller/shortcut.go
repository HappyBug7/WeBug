package controller

import (
	"back_end/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Shortcut struct {
}

func (s *Shortcut) Link_add(c *gin.Context) {
	var CheckInfo struct {
		Link  string `json:"link" form:"link"`
		Title string `json:"title" form:"title"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	err := srv.Shortcut.Link_add(CheckInfo.Link, CheckInfo.Title)
	c.JSON(http.StatusOK, ResponseNew(c, err))
}

func (s *Shortcut) Links_get(c *gin.Context) {
	err, resp := srv.Shortcut.Links_get()
	if err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
