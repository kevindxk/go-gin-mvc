/**
 * @Author: dexukong
 * @Description:
 * @File:  problemController
 * @Version: 1.0.0
 * @Date: 2022/08/09 16:36
 */

package controller

import (
	"github.com/gin-gonic/gin"
	"goproject/admin/services"
)

type ProblemController struct {
}

func (problem ProblemController) GetList(c *gin.Context) {
	services.ProblemServices{}.GetList(c)
}

func (problem ProblemController) GreatProblem(c *gin.Context) {
	services.ProblemServices{}.GreatProblem(c)
}
