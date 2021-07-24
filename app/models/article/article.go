package article

import (
	"goblog/app/models"
	"goblog/pkg/route"
	"goblog/pkg/types"
)

type Article struct {
	models.BaseModel

	Title string
	Body  string
}

// Link 方法用来生成文章链接
func (a Article) Link() string  {
	return route.Name2URL("articles.show", "id", a.GetStringID())
}

func (a Article) GetStringID() string {
	return types.Uint64ToString(a.ID)
}