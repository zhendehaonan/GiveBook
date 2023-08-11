package controller

import (
	"GiveBook/common/dao"
	"GiveBook/model"
	"GiveBook/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
)

func Register(c *gin.Context) {
	DB := dao.GetDB()
	//获取参数
	name := c.PostForm("name")
	telephone := c.PostForm("telephone")
	password := c.PostForm("password")
	//做数据验证
	if len(telephone) != 11 {
		c.JSON(422, "手机号必须为11位")
		return
	}
	if len(password) < 6 {
		c.JSON(422, "密码不能少于6位")
		return
	}
	//如果名字没有传，系统会自动给一个10位的字符串
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name)
	log.Println(telephone)
	log.Println(password)
	//判断手机号是否存在
	//手机号已存在
	if isTelephoneExist(DB, telephone) {
		c.JSON(422, "用户已经存在，不允许注册")
		return
	}
	//手机号不存在，则新建用户
	dao.InitDB().Create(&model.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	})

	c.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

// 判断手机号是否已经存在
func isTelephoneExist(DB *gorm.DB, telephone string) bool {
	var user model.User
	DB.Where("telephone=?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
