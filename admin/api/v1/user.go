/**
 * @Author: dexukong
 * @Description:
 * @File:  user
 * @Version: 1.0.0
 * @Date: 2022/08/23 10:04
 */

package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goproject/admin/services"
	"net/http"
)

func UserRegister(c *gin.Context) {
	var userRegister services.UserService
	fmt.Println(c.ShouldBindJSON(&userRegister))
	if err := c.ShouldBindJSON(&userRegister); err != nil {
		res := userRegister.Register(c.Request.Context())
		c.JSON(http.StatusOK, res)
	} else {
		c.JSON(http.StatusOK, err)
	}
}
