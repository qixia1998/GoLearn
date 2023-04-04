package beego

import (
	"github.com/beego/beego/v2/server/web"
	"testing"
)

func TestUserController(t *testing.T) {
	web.BConfig.CopyRequestBody = true
	c := &UserController{}
	web.Router("/user", c, "get:GetUser")
	// 监听 8081 端口
	web.Run(":8081")
}
