package controller

import (
	"first/go_web/common"
	"first/go_web/model"
	"first/go_web/util"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
)

func Register(ctx  *gin.Context) {
	DB := common.GetDB()
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")
	if len(telephone)!=11{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号必须为11位"})
		return
	}
	if len(password)<6{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}
	if len(password)<6{
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
		return
	}
	if len(name) ==0{
		name=util.RandomString(10)
	}
	log.Println(name,telephone,password)
	if isTelephoneExist(DB,telephone){
		ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"用户已经存在"})
		return
	}

	newUser := model.User{
		Name:name,
		Telephone: telephone,
		Password: password,
	}
	DB.Create(&newUser)
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelephoneExist(db *gorm.DB,telephone string)bool{
	var user model.User
	db.Where("telephone = ?",telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}