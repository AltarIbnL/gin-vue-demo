package controller

import (
	"github.com/gin-gonic/gin"
	"myproject/Model"
	"myproject/repository"
	"myproject/response"
	"myproject/vo"
	"strconv"
)

//只要带有 ResController 里面4个方法的都可以赋值为ICategoryController这个接口
type ICategoryController interface {
	ResController
}

type CatehoryController struct {
	//DB *gorm.DB
	Repository repository.CategoryRepository
}

func NewCategoryController() ICategoryController {
	repository := repository.NewCategoryRepository()
	repository.DB.AutoMigrate(&Model.Category{})
	return CatehoryController{Repository: repository}
}

func (c2 CatehoryController) Create(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	//自动绑定传入的类型，Bind 返回的是文本的形式，shouldBind返回的是JSON格式的
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}
	category, err := c2.Repository.Create(requestCategory.Name)
	if err != nil {
		panic(err)
		response.Fail(c, nil, "创建失败")
		return
	}
	response.Success(c, gin.H{"category": category}, "")
}

func (c2 CatehoryController) Update(c *gin.Context) {
	var requestCategory vo.CreateCategoryRequest
	//自动绑定传入的类型,Bind 返回的是文本的形式，shouldBind返回的是JSON格式的
	if err := c.ShouldBind(&requestCategory); err != nil {
		response.Fail(c, nil, "数据验证错误，分类名称必填")
		return
	}

	//获取path的参数
	catagoryID, _ := strconv.Atoi(c.Params.ByName("id"))
	//catagoryID := requestCategory.ID

	updateCategory, err := c2.Repository.SelectByID(catagoryID)
	if err != nil {
		response.Fail(c, nil, "分类不存在")
		return
	}

	//更新分类
	category, err := c2.Repository.Update(*updateCategory, requestCategory.Name)
	if err != nil {
		panic(err)
	}
	response.Success(c, gin.H{"category": category}, "修改成功")

}

func (c2 CatehoryController) Show(c *gin.Context) {
	//获取path的参数
	catagoryID, _ := strconv.Atoi(c.Params.ByName("id"))
	//catagoryID := requestCategory.ID
	//在数据库中查找ID相同的
	category, err := c2.Repository.SelectByID(catagoryID)
	if err != nil { //如果没找到
		response.Fail(c, nil, "分类不存在")
		return
	}
	response.Success(c, gin.H{"category": category}, "")

}

func (c2 CatehoryController) Delete(c *gin.Context) {
	//获取path的参数
	catagoryID, _ := strconv.Atoi(c.Params.ByName("id"))
	//在数据库中查找ID相同的
	if err := c2.Repository.DeleteById(catagoryID); err != nil { //如果没找到
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")

}
