package requests

import (
	"github.com/thedevsaddam/govalidator"
	"goblog/app/models/user"
)

// ValidateRegistrationForm 验证表单，返回 errs 长度等于零即可
func ValidateRegistrationForm(data user.User) map[string][]string {
	// 1. 定制认证规则
	rules := govalidator.MapData{
		"name":             []string{"required", "alpha_num", "between:3,20"},
		"email":            []string{"required", "min:4", "max:30", "email"},
		"password":         []string{"required", "min:6"},
		"password_confirm": []string{"required"},
	}

	// 2. 定制错误消息
	messages := govalidator.MapData{
		"name": []string{
			"required:Name is required",
			"alpha_num:The format is wrong, only numbers and English are allowed",
			"between:Name length must be between 3~20",
		},
		"email": []string{
			"required:Email is required",
			"min:Email length must be greater than 4",
			"max:Email length must be less than 30",
			"email:The email format is incorrect, please provide a valid email address",
		},
		"password": []string{
			"required:Password is required",
			"min:Password length must be greater than 6",
		},
		"password_confirm": []string{
			"required:Password confirm is required",
		},
	}

	// 3. 配置选项
	opts := govalidator.Options{
		Data:          &data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	// 4. 开始认证
	errs := govalidator.New(opts).ValidateStruct()

	// 5. 因 govalidator 不支持 password_confirm 验证，我们自己写一个
	if data.Password != data.PasswordConfirm {
		errs["password_confirm"] = append(errs["password_confirm"], "The passwords entered twice do not match!")
	}

	return errs
}
