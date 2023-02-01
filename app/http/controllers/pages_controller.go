package controllers

import (
	"fmt"
	"net/http"
)

// PagesController 处理静态页面
type PagesController struct {
}

// Home 首页
func (*PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Hello, Here it is goblog!</h1>")
}

// About 关于我们页面
func (*PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "About me")
}

// NotFound 404 页面
func (*PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "404 Page not found")
}
