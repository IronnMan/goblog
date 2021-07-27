package view

import (
	"goblog/pkg/auth"
	"goblog/pkg/logger"
	"goblog/pkg/route"
	"html/template"
	"io"
	"path/filepath"
	"strings"
)

type D map[string]interface{}

// Render 渲染通用视图
func Render(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "app", data, tplFiles...)
}

// RenderSimple 渲染简单的视图
func RenderSimple(w io.Writer, data D, tplFiles ...string) {
	RenderTemplate(w, "simple", data, tplFiles...)
}

// RenderTemplate 渲染视图
func RenderTemplate(w io.Writer, name string, data D, tplFiles ...string)  {

	// 通用模版数据
	data["isLogined"] = auth.Check()

	// 生成模版文件
	allFiles := getTemplateFiles(tplFiles...)

	// 解析所有模版文件
	tmpl, err := template.New("").Funcs(template.FuncMap{
	    	"RouteName2URL": route.Name2URL,
	    }).ParseFiles(allFiles...)
	logger.LogError(err)

	// 渲染模版
	tmpl.ExecuteTemplate(w, name, data)
}

func getTemplateFiles(tplFiles ...string) []string {
	// 设置模版相对路径
	viewDir := "resources/views/"

	// 遍历传参文件列表 Slice，设置正确的路径，支持 dir.filename 语法糖
	for i, f := range tplFiles {
		tplFiles[i] = viewDir + strings.Replace(f, ".", "/", -1) + ".gohtml"
	}

	// 所有布局模版文件 Slice
	layoutFiles, err := filepath.Glob(viewDir + "layouts/*.gohtml")
	logger.LogError(err)

	return append(layoutFiles, tplFiles...)
}