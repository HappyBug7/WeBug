package controller

import (
	"back_end/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Search struct {
}

func (s *Search) Suggestion_get(c *gin.Context) {
	var CheckInfo struct {
		Q string `json:"q" form:"q"`
	}
	if err := c.ShouldBind(&CheckInfo); err != nil {
		c.Error(common.ErrNew(err, common.SysErr))
		return
	}
	err, data := srv.Search.Suggestion_get(CheckInfo.Q)
	resp := struct {
		Err  error
		Data interface{}
	}{
		Err:  err,
		Data: data,
	}
	c.JSON(http.StatusOK, ResponseNew(c, resp))
}
