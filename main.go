package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.Run()
}

//CreateOwner 新建立车主账号到 车主信息数据库
func CreateOwner(c *gin.Context) {

}

//QueryOwner 查询车主账号是否在 车主信息数据库
func QueryOwner(c *gin.Context) {

}
