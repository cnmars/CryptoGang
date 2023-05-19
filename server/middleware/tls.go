package middleware



import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"fmt"
	"golddigger/global"
)

/**
*添加对https的支持
 用https把这个中间件在router里面use一下就好
 */
func TlsHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     global.GD_CONFIG.SslConfig.Domain,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			// 如果出现错误，请不要继续
			fmt.Println(err)
			return
		}
		// 继续往下处理
		c.Next()
	}
}


