package Middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wzhyyds123/golibrary/jwt"
	"github.com/wzhyyds123/golibrary/log"
	"strings"
)

func AuthMiddleware(c *gin.Context){
	token := c.Request.Header.Get("Authorization")
	if token == "" {
		c.JSON(401,gin.H{
			"msg":  "权限不足",
			"code": 401,
		})
		log.Error.Println("权限不足")
		panic("权限不足")
	}else {	kv := strings.Split(token, " ")
		if len(kv) != 2 || kv[0] != "Bearer" {
			panic("AuthString无效")
		}
		tokenString := kv[1]
		_,err:=jwt.ValidateToken(tokenString)
		fmt.Println(err)
		if err!=nil {
			c.JSON(401,gin.H{
				"msg":  "权限不足",
				"code": 401,
			})
			log.Error.Println("权限不足")
			panic("权限不足")
		}}
	}