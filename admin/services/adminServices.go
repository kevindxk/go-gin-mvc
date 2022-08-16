/**
 * @Author: dexukong
 * @Description:
 * @File:  adminServices
 * @Version: 1.0.0
 * @Date: 2022/08/08 10:12
 */

package services

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminServices struct {
}

//RegisterService
// @BasePath /admin/register
// @Tags 公共方法
// @Summary 用户注册
// @Param data body define.UserInfo true "UserInfo"
// @Param email formData string true "email"
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /register [post]
func (adminService AdminServices) RegisterService(c *gin.Context) {

	//paramData := new(define.UserInfo)
	//err := c.ShouldBindJSON(paramData)
	//if err != nil {
	//	fmt.Printf(paramData.UserName, paramData.Password, paramData.Email, paramData.Phone)
	//}
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")
	email := c.PostForm("email")
	//data := db.First(&model.UserLogin{})
	//userJson := model.UserLogin{}
	//c.BindJSON(&userJson)
	//c.JSON(http.StatusOK, gin.H{
	//	"admin/login": "这是后台用户登录接口",
	//	"username":    userJson.UserName,
	//	"password":    userJson.Password,
	//	"phone":       userJson.Phone,
	//	"emain":       userJson.Email,
	//	"data":        data,
	//})
	//json := model.User{}
	//c.BindJSON(&json)
	//username := json.UserName
	//password := json.Password
	//phone := json.Phone
	//email := json.Email
	if email == "" || username == "" || password == "" || phone == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": "-1",
			"msg":  "参数不对",
		})
		return
	}
	//db := utils.Init()
	//var cnt int64
	//err = db.Where("email=?", email).Model(new(model.User)).Count(&cnt).Error
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": "-1",
	//		"msg":  "Get User Error:" + err.Error(),
	//	})
	//	return
	//}
	//if cnt > 0 {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  "该邮箱已被注册",
	//	})
	//	return
	//}
	////数据写入
	//userIdentity := utils.GetUUID()
	//data := &model.User{
	//	Identity: userIdentity,
	//	UserName: username,
	//	Password: utils.GetMd5(password),
	//	Phone:    phone,
	//	Email:    email,
	//}
	//err = db.Create(data).Error
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  "Crete User Error:" + err.Error(),
	//	})
	//	return
	//}
	////生成Token
	//token, err := utils.GenerateToken(userIdentity, username, data.IsAdmin)
	//if err != nil {
	//	c.JSON(http.StatusOK, gin.H{
	//		"code": -1,
	//		"msg":  "Generate Token Error:" + err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"code": 200,
	//	"data": map[string]interface{}{
	//		"token": token,
	//	},
	//})
	c.JSON(http.StatusOK, gin.H{
		"username": username,
		"password": password,
		"phone":    phone,
		"emain":    email,
	})
}
