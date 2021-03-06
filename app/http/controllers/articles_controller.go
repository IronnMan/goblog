package controllers

import (
	"fmt"
	"goblog/app/models/article"
	"goblog/app/policies"
	"goblog/app/requests"
	"goblog/pkg/auth"
	"goblog/pkg/route"
	"goblog/pkg/view"
	"net/http"
)

// ArticlesController 文章相关页面
type ArticlesController struct {
	BaseController
}


// Store 文章创建页面
func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {

	// 初始化数据
	currentUser := auth.User()
	_article := article.Article{
		Title: r.PostFormValue("title"),
		Body:  r.PostFormValue("body"),
		UserID: currentUser.ID,
	}

	errors := requests.ValidateArticleForm(_article)

	// 检查是否有错误
	if len(errors) == 0 {
		// 创建文章
		_article.Create()
		if _article.ID > 0 {
			indexURL := route.Name2URL("articles.show", "id", _article.GetStringID())
			http.Redirect(w, r, indexURL, http.StatusFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Article": _article,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}


// Create 文章创建页面
func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request)  {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

// Show 文章详情页面
func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	// 获取 URL 参数
	id := route.GetRouteVariable("id", r)

	// 读取对应的文章数据
	article, err := article.Get(id)

	// 如果出现错误
	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {

		// 读取成功，显示文章
		view.Render(w, view.D{
			"Article": article,
			"CanModifyArticle": policies.CanModifyArticle(article),
		}, "articles.show", "articles._article_meta")
	}
}

// Index 文章列表页
func (ac *ArticlesController) Index(w http.ResponseWriter, r *http.Request) {

	// 获取结果集
	articles, pagerData, err := article.GetAll(r, 5)

	if err != nil {
		// 数据库错误
		ac.ResponseForSQLError(w, err)
	} else {
		// 加载模版
		view.Render(w, view.D{
			"Articles": articles,
			"PagerData": pagerData,
		}, "articles.index", "articles._article_meta")
	}
}

// Edit 文章更新页面
func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		// 检查权限
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			// 读取成功，显示比编辑文章表单
			view.Render(w, view.D{
				"Article": _article,
				"Errors": view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}
}

// Update 更新文章
func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {

	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {

		// 检查权限
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			_article.Title = r.PostFormValue("title")
			_article.Body  = r.PostFormValue("body")

			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {

				rowsAffected, err := _article.Update()

				if err != nil {
					// 数据库错误
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
					return
				}

				// 更新成功，跳转到文章详情页
				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", id)
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改！")
				}
			} else {

				// 表单验证不通过，显示理由
				view.Render(w, view.D{
					"Article": _article,
					"Errors": errors,
				}, "articles.edit", "articles._form_field")
			}
		}
	}
}

// Delete 删除文章
func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request)  {

	id := route.GetRouteVariable("id", r)

	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {

		// 检查权限
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			rowsAffected, err := _article.Delete()

			if err != nil {
				// SQL 错误
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				if rowsAffected > 0 {
					indexURL := route.Name2URL("articles.index")
					http.Redirect(w, r, indexURL, http.StatusNotFound)
				} else {
					// Edge casae
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}