package controllers

import (
	"fmt"
	"goblog/app/models/category"
	"goblog/app/requests"
	"goblog/pkg/flash"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// CategoriesController 文章分类控制器
type CategoriesController struct {
	BaseController
}

// Create 文章分类创建页面
func (*CategoriesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "categories.create")
}

// Store 保存文章分类
func (*CategoriesController) Store(w http.ResponseWriter, r *http.Request) {
	// 1. 初始化数据
	_category := category.Category{
		Name: r.PostFormValue("name"),
	}

	// 2. 表单验证
	errors := requests.ValidateCategoryForm(_category)

	// 3. 检测错误
	if len(errors) == 0 {
		// 创建文章分类
		_category.Create()
		if _category.ID > 0 {
			flash.Success("Created successfully!")
			indexURL := route.Name2URL("home")
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "Failed to create article categories, please contact the administrator")
		}
	} else {
		view.Render(w, view.D{
			"Category": _category,
			"Errors":   errors,
		}, "categories.create")
	}
}

// Show 显示分类下的文章列表
func (*CategoriesController) Show(w http.ResponseWriter, r *http.Request) {
	//
}
