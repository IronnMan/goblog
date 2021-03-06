package controllers

import (
	"goblog/app/models/user"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/flash"
	"goblog/pkg/view"
	"net/http"
)

// AuthController 处理静态页面
type AuthController struct {

}

// userForm 用户表单
type userForm struct {
	Name            string `valid:"name"`
	Email           string `valid:"email"`
	Password        string `valid:"password"`
	PasswordConfirm string `valid:"password_confirm"`
}

// Register 注册页面
func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

// DoRegister 处理注册逻辑
func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	// 0. 初始化数据
	_user := user.User{
		Name:            r.PostFormValue("name"),
		Email:           r.PostFormValue("email"),
		Password:        r.PostFormValue("password"),
		PasswordConfirm: r.PostFormValue("password_confirm"),
	}

	// 1. 表单规则
	errs := requests.ValidateRegistrationForm(_user)


	if len(errs) > 0 {
		// 有错误发生，打印数据
		view.RenderSimple(w, view.D{
			"Errors": errs,
			"User":   _user,
		}, "auth.register")


	} else {
		// 验证通过，创建数据
		_user.Create()

		if _user.ID > 0 {
			// 登陆用户并跳转到首页
			flash.Success("恭喜您注册成功！")
			auth.Login(_user)
			http.Redirect(w, r, "/", http.StatusFound)
		}
	}
	// 3. 表单不通过 -- 重新显示表单
}


// Login 显示登陆表单
func (*AuthController) Login(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.login")
}

// DoLogin 处理登陆表单提交
func (*AuthController) DoLogin(w http.ResponseWriter, r *http.Request) {

	// 初始化表单数据
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	// 尝试登陆
	if err := auth.Attempt(email, password); err == nil {
		// 登陆成功
		flash.Success("欢迎回来！")
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		// 失败，显示错误提示
		view.RenderSimple(w, view.D{
			"Error": err.Error(),
			"Email": email,
			"Password": password,
		}, "auth.login")
	}
}

// Logout 退出登陆
func (*AuthController) Logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	http.Redirect(w, r, "/", http.StatusFound)
}

// Logout 退出登陆
func (*AuthController) logout(w http.ResponseWriter, r *http.Request) {
	auth.Logout()
	flash.Success("您已退出登陆")
	http.Redirect(w, r, "/", http.StatusFound)
}