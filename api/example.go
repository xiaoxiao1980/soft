package api

import (
	"app/global/resp"

	"github.com/gin-gonic/gin"
)

// @Summary 监控
// @Description 获取系统信息
// @Tags default
// @Success 200 {string} string "{"message": "ok"}"
// @Router /monitor/server [get]
// @Security Bearer
func funcname(c *gin.Context) {

	resp.Ok(c)

}
