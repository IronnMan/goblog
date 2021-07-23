package view

import (
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

// Render 渲染视图
func Render(w io.Writer, name string, data interface{})  {

	// 设置模版相对路径
	viewDir := "resources/views/"

	// 语法糖，将 articles.show 更早为 articles/show
	name = strings.Replace(name, ".", "/", -1)

	// 所有布局模版文件 Slice
	files, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	// 在 Slice 里新增我们的目标文件
	newFiles := append(files, viewDir + name + ".gohtml")

	// 解析所有模版文件
	tmpl, err := template.New(name + ".gohtml").Funcs(template.FuncMap{
	    	"RouteName2URL": route.Name2URL,
	    }).ParseFiles(newFiles...)
	logger.LogError(err)

	// 渲染模版
	tmpl.ExecuteTemplate(w, "app", data)
}