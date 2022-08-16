/**
 * @Author: dexukong
 * @Description:
 * @File:  adminUser
 * @Version: 1.0.0
 * @Date: 2022/08/08 10:08
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"goproject/admin/services"
)

type AdminController struct {
}

func (admin AdminController) Register(c *gin.Context) {

	services.AdminServices{}.RegisterService(c)
}
