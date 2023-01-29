package tests

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestHomePage(t *testing.T) {
	baseURL := "http://localhost:3000"

	// 1. 请求 -- 模拟用户访问浏览器
	var (
		resp *http.Response
		err  error
	)
	resp, err = http.Get(baseURL + "/")

	// 2. 检测 -- 是否无错且 200
	assert.NoError(t, err, "An error occurred, err is not empty!")
	assert.Equal(t, 200, resp.StatusCode, "Should return status code 200")
}
