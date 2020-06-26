package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"redrock0309/token"
	"redrock0309/resps"
)


func User(c *gin.Context) {
	auth:= c.GetHeader("token")
	fmt.Println(auth)
	if len(auth)<7 {
		resps.Error(c, 10011, "token error")
		c.Abort()//阻止调用后续处理函数
		return
	}
	
	uid, username, err := token.CheckToken(auth)
	fmt.Println(err)
	if err != nil {
		resps.Error(c, 10011, "token error")
		c.Abort()
		return
	}
	c.Set("id", uid)
	c.Set("username", username)//类型数据结构的初始化操作
	c.Next()//调用后续处理函数
	return
}