package v1

import (
	"axiangcoding/antonstar/api-system/internal/app/entity"
	"axiangcoding/antonstar/api-system/internal/app/service"
	"axiangcoding/antonstar/api-system/pkg/app"
	"axiangcoding/antonstar/api-system/pkg/app/e"
	"github.com/gin-gonic/gin"
)

// LoginForm 目前支持用户名登录
type LoginForm struct {
	UserName string `binding:"username,required" json:"username" form:"username"`
	Password string `binding:"password,required" json:"password" form:"password"`
}

// UserLogin
// @Summary  User login
// @Tags      User
// @Param    form  body      LoginForm    true  "register form"
// @Success  200   {object}  app.ApiJson  ""
// @Failure  400   {object}  app.ErrJson  ""
// @Router   /v1/user/login [post]
func UserLogin(c *gin.Context) {
	form := LoginForm{}
	err := c.ShouldBindJSON(&form)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}

	login := entity.UserLogin{
		UserName: form.UserName,
		Password: form.Password,
	}

	token, err := service.UserLogin(c, login)
	if err != nil {
		app.BizFailed(c, e.LoginFailed, err)
		return
	}
	app.Success(c, map[string]string{app.AuthHeader: token})
}

type RegisterForm struct {
	UserName  string `binding:"username,required" json:"username" form:"username"`
	Email     string `binding:"email,required" json:"email" form:"email"`
	Phone     string `binding:"omitempty,e164" json:"phone"`
	AvatarUrl string `binding:"omitempty,url" json:"avatar_url"`
	Password  string `binding:"password,required" json:"password" form:"password"`
}

// UserRegister
// @Summary  用户注册
// @Tags     User
// @Param    form  body      RegisterForm  true  "register form"
// @Success  200  {object}  app.ApiJson  ""
// @Failure  400   {object}  app.ErrJson   ""
// @Router   /v1/user/register [post]
func UserRegister(c *gin.Context) {
	regForm := RegisterForm{}
	err := c.ShouldBindJSON(&regForm)
	if err != nil {
		app.BadRequest(c, e.RequestParamsNotValid, err)
		return
	}
	register := entity.UserRegister{
		UserName: regForm.UserName,
		Email:    regForm.Email,
		Phone:    regForm.Phone,
		Password: regForm.Password,
	}
	id, err := service.UserRegister(c, register)
	if err != nil {
		app.BizFailed(c, e.RegisterFailed, err)
		return
	}
	app.Success(c, map[string]string{"id": id})
}

// UserLogout
// @Summary   User logout
// @Tags     User
// @Success   200  {object}  app.ApiJson  ""
// @Failure   400  {object}  app.ErrJson  ""
// @Router    /v1/user/logout [post]
// @Security  ApiKeyAuth
func UserLogout(c *gin.Context) {
	result, err := service.UserLogout(c, c.GetHeader(app.AuthHeader))
	if err != nil {
		app.BizFailed(c, e.Error, err)
		return
	}
	if result == 0 {
		app.BizFailed(c, e.TokenNotValid)
		return
	}
	app.Success(c, nil)
}

type KeyFieldExistForm struct {
}

// IsKeyFieldValueExist
// @Summary  判断主要的用户信息的值是否存在
// @Tags     User
// @Success  200   {object}  app.ApiJson   ""
// @Failure  400  {object}  app.ErrJson  ""
// @Router   /v1/user/key-field/exist [post]
func IsKeyFieldValueExist(c *gin.Context) {
	// TODO
	app.Success(c, nil)
}
