package middleware

import (
	"github.com/gin-gonic/gin"
)

/*
*
域名中间件
*/
func DomainLimitMiddleware(c *gin.Context) {
	//离线或者远程
	if !CheckBindOffcial(c) {
		c.Abort()
		return
	}

}

// 绑定官网账户
func CheckBindOffcial(c *gin.Context) bool {
	//host := c.Request.Host
	//fmt.Println(host)
	//res, err := tools.HTTPGet("https://gofly.v1kf.com/2/isBindOfficial")
	//if err != nil {
	//	log.Println("离线授权码失败,认证连接失败")
	//	c.Redirect(302, "/bind")
	//	c.Abort()
	//}
	//if string(res) != "success" {
	//	c.Redirect(302, "/bind")
	//	c.Abort()
	//}
	return true
}
