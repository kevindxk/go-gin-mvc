/**
 * @Author: dexukong
 * @Description:
 * @File:  router
 * @Version: 1.0.0
 * @Date: 2022/08/08 10:00
 */

package router

import (
	"github.com/gin-gonic/gin"
	"goproject/admin/controller"
	"net/http"
)

func RouterInit(r *gin.Engine) {

	adminRouter := r.Group("/admin")

	adminRouter.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"admin": "欢迎来到Admin后台",
		})
	})

	adminRouter.POST("/register", controller.AdminController{}.Register)
	adminRouter.GET("/problemList", controller.ProblemController{}.GetList)
	adminRouter.POST("/problemCreat", controller.ProblemController{}.GreatProblem)
}
