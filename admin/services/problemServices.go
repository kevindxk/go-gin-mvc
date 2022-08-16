/**
 * @Author: dexukong
 * @Description:
 * @File:  problemServices
 * @Version: 1.0.0
 * @Date: 2022/08/09 16:39
 */

package services

import (
	"github.com/gin-gonic/gin"
	"goproject/admin/define"
	"goproject/admin/model"
	"goproject/admin/utils"
	"log"
	"net/http"
	"strconv"
)

type ProblemServices struct {
}

//RegisterService
// @BasePath /admin/problemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problemList [get]
func (problemServices ProblemServices) GetList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaulPage))
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaulSize))
	if err != nil {
		log.Panicln("getList page strconv Error", err)
		return
	}
	page = (page - 1) * size
	var count int64
	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")
	list := make([]*model.ProblemBasic, 0)
	tx := utils.Init().Model(new(model.ProblemBasic)).Distinct("`problem_basic`.`id`").Select("DISTINCT(`problem_basic`.`id`), `problem_basic`.`identity`, "+
		"`problem_basic`.`title`, `problem_basic`.`max_runtime`, `problem_basic`.`max_mem`, `problem_basic`.`pass_num`, "+
		"`problem_basic`.`submit_num`, `problem_basic`.`created_at`, `problem_basic`.`updated_at`, `problem_basic`.`deleted_at` ").Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? OR content like ? ", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id FROM category_basic cb WHERE cb.identity = ? )", categoryIdentity)
	}
	err = tx.Count(&count).Error
	if err != nil {
		log.Println("GetProblemList Count Error:", err)
		return
	}
	err = tx.Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		log.Println("GetProblemList Count Error:", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": "200",
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}

// ProblemCreate
// @Tags 管理员私有方法
// @Summary 问题创建
// @Accept json
// @Param authorization header string true "authorization"
// @Param data body define.ProblemBasic true "ProblemBasic"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problemCreat [post]
func (problemServices ProblemServices) GreatProblem(c *gin.Context) {
	in := new(define.ProblemBasic)
	err := c.ShouldBindJSON(in)
	if err != nil {
		log.Println("[JsonBind Error] : ", err)
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数错误",
		})
		return
	}
	if in.Title == "" || in.Content == "" || len(in.ProblemCategories) == 0 || len(in.TestCases) == 0 || in.MaxRuntime == 0 || in.MaxMem == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "参数不能为空",
		})
		return
	}
	identity := utils.GetUUID()

	data := &model.ProblemBasic{
		Identity:   identity,
		Title:      in.Title,
		Content:    in.Content,
		MaxRuntime: in.MaxRuntime,
		MaxMem:     in.MaxMem,
	}
	// 处理分类
	categoryBasics := make([]*model.ProblemCategory, 0)
	for _, id := range in.ProblemCategories {
		categoryBasics = append(categoryBasics, &model.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(id),
		})
	}
	data.ProblemCategories = categoryBasics

	// 处理测试用例
	testCaseBasics := make([]*model.TestCase, 0)
	for _, v := range in.TestCases {
		// 举个例子 {"input":"1 2\n","output":"3\n"}
		testCaseBasic := &model.TestCase{
			Identity:        utils.GetUUID(),
			ProblemIdentity: identity,
			Input:           v.Input,
			Output:          v.Output,
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics
	err = utils.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "Problem Create Error:" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}
