package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/article"
)

func ValidateArticleForm(data article.Article) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"title": []string{"required", "min:3", "max:40"},
		"body":  []string{"required", "min:10"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"title": []string{
			"required:Title is required",
			"min:Title length must be greater than 3",
			"max:Title length must be less than 40",
		},
		"body": []string{
			"required:Body is required",
			"min:Body length must be greater than 10",
		},
	}

	// 3. 配置初始化
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// 4. 开始验证
	return govalidator.New(opts).ValidateStruct()
}