/**
 * @Author: dexukong
 * @Description:
 * @File:  main
 * @Version: 1.0.0
 * @Date: 2022/08/08 09:56
 */

package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"goproject/admin/main/docs"
	"goproject/admin/router"
)

func main() {
	r := gin.Default()

	router.RouterInit(r)
	//r.GET("/admin", func(context *gin.Context) {
	//	context.JSON(http.StatusOK, gin.H{
	//		"admin": "这是admin后台",
	//	})
	//})
	// swagger访问地址 http://localhost:8080/swagger/index.html
	//配置swagger访问地址
	docs.SwaggerInfo.BasePath = "/admin"
	url := ginSwagger.URL("http://localhost:9999/swagger/doc.json")

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	r.Run(":9999")
}
