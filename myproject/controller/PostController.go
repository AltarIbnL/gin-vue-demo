package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"myproject/Model"
	"myproject/common"
	"myproject/response"
	"myproject/vo"
	"strconv"
)

type IPostController interface {
	ResController
	PageList(c *gin.Context)
}
type PostController struct {
	DB *gorm.DB
}

func (p PostController) Create(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	// 数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Print(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	// 获取登录用户
	user, _ := c.Get("user")

	// 创建post
	post := Model.Post{
		UserId:     user.(Model.User).ID,
		CategoryId: requestPost.CategoryId,
		Title:      requestPost.Title,
		HeadImg:    requestPost.HeadImg,
		Content:    requestPost.Content,
	}

	// 插入数据
	if err := p.DB.Create(&post).Error; err != nil {
		panic(err)
		return
	}

	// 成功
	response.Success(c, nil, "创建成功")

}

func (p PostController) Update(c *gin.Context) {
	var requestPost vo.CreatePostRequest
	//数据验证
	if err := c.ShouldBind(&requestPost); err != nil {
		log.Println(err.Error())
		response.Fail(c, nil, "数据验证错误")
		return
	}

	//获取path
	postID := c.Params.ByName("id")
	var post Model.Post
	//如果不存在
	if p.DB.Where("id=?", postID).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}

	// 获取登陆用户
	user, _ := c.Get("user")
	// 判断当前用户是否为文章的作者
	userId := user.(Model.User).ID
	if userId != post.UserId {
		response.Fail(c, nil, "文章匹配错误")
		return
	}
	//更新文章
	if err := p.DB.Model(&post).Update(requestPost).Error; err != nil {
		response.Fail(c, nil, "更新失败")
		return
	}
	response.Success(c, gin.H{"post": post}, "更新成功")
}

func (p PostController) Show(c *gin.Context) {
	//获取path
	postID := c.Params.ByName("id")
	var post Model.Post
	//如果不存在
	if p.DB.Preload("Category").Where("id=?", postID).First(&post).RecordNotFound() {
		response.Fail(c, nil, "文章不存在")
		return
	}
	response.Success(c, gin.H{"post": post}, "查询成功")

}

func (p PostController) Delete(ctx *gin.Context) {
	// 获取分页参数
	pageNum, _ := strconv.Atoi(ctx.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", "20"))

	// 分页
	var posts []Model.Post
	p.DB.Order("created_at desc").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 记录的总条数
	var total int
	p.DB.Model(Model.Post{}).Count(&total)

	// 返回数据
	response.Success(ctx, gin.H{"data": posts, "total": total}, "成功")

}

func (p PostController) PageList(c *gin.Context) {
	// 获取分页参数
	PageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "20"))
	//分页
	var posts []Model.Post
	//按照创建时间进行分页
	p.DB.Order("created_at desc").Offset((PageNum - 1) * pageSize).Limit(pageSize).Find(&posts)

	// 前段渲染分页需要知道的综述
	var total int
	p.DB.Model(&Model.Post{}).Count(&total)

	response.Success(c, gin.H{"data": posts, "total": total}, "成功")

}

func NewPostController() IPostController {
	db := common.GetDB()
	db.AutoMigrate(&Model.Post{})
	return PostController{DB: db}
}
