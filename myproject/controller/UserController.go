package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"myproject/Model"
	"myproject/common"
	"myproject/dto"
	"myproject/response"
	"myproject/util"
	"net/http"
)

//var Db *gorm.DB
func Register(c *gin.Context) {
	Db := common.GetDB()
	//获取参数
	//name := c.PostForm("name") //在网页端以关键字名称的形式输入
	//telephone := c.PostForm("telephone")
	//passwd := c.PostForm("passwd")
	var u1 Model.User
	err := c.ShouldBind(&u1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	name := u1.Name
	telephone := u1.Telephone
	passwd := u1.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号必须为11位",
		//})
		return
	}
	if len(passwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "密码不能少于6位",
		//})
		return
	}

	//如果名称没有传入，给一个10位的随机字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}

	log.Println(name, telephone, passwd)
	//判断手机号是否存在
	if isTelephoneExist(Db, telephone) {
		//不允许注册
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户已经存在")

		return
	}
	//fmt.Println(passwd)
	//创建用户
	hashpass, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "加密错误")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 500,
		//	"msg":  "加密错误",
		//})
		return
	}
	newUser := Model.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashpass),
	}
	//fmt.Println(newUser)
	Db.Create(&newUser)

	token, err := common.ReleaseToken(newUser)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generated error:%v", err)
		return
	}
	//返回结果
	//response.Success(c, nil, "注册成功")
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"msg":   "注册成功",
		"token": token,
	})
	//发放token

}

func Login(c *gin.Context) {
	db := common.GetDB()
	//获取参数
	var u1 Model.User
	err := c.ShouldBind(&u1)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	telephone := u1.Telephone
	passwd := u1.Password

	// 数据验证
	if len(telephone) != 11 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "手机号必须为11位")
		//c.JSON(http.StatusUnprocessableEntity, gin.H{
		//	"code": 422,
		//	"msg":  "手机号必须为11位",
		//})
		return
	}
	if len(passwd) < 6 {
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "密码不能少于6位位")
		return
	}

	//判断手机号是否存在
	var user Model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID == 0 {
		//不存在
		response.Response(c, http.StatusUnprocessableEntity, 422, nil, "用户不存在")

		return
	}

	//判断密码是否正确,将传入的和数据库中用户的密码比较
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(passwd)); err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 400, nil, "密码错误")
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(c, http.StatusUnprocessableEntity, 500, nil, "系统异常")
		log.Printf("token generated error:%v", err)
		return
	}
	//返回结果
	c.JSON(http.StatusOK, gin.H{
		"code":  200,
		"token": token,
		"msg":   "登陆成功",
	})

}

func isTelephoneExist(db *gorm.DB, telephone string) bool {
	var user Model.User
	db.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}

func Info(c *gin.Context) {
	// 在AuthMiddleware() 中间件中将 user上传上去了
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": dto.ToUserDto(user.(Model.User))}, //类型断言，user是一个被赋值为Model.User的interface{}, 此时我们
	})

}
