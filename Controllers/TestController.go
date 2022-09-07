package Controllers

import (
	"github.com/gin-gonic/gin"
	"go-gin/Models"
	"net/http"
)

func WzhFind(c *gin.Context) {
	active_name := c.PostForm("active_name")
	res, err := Models.WzhFind(active_name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    100,
			"message": "查询失败",
			"data":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "查询成功",
			"data":    res,
		})
	}
}