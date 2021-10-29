// @Author: float311@163.com
// @Description:
// @Version: 1.0.0
// Date: 2021/10/25 21:31
// Update: 2021/10/25 21:31

package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/lzwgiter/Go-Blog/internal/routers/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	apiV1 := r.Group("/api/v1")

	article := v1.NewArticle()
	tag := v1.NewTag()

	{
		// 标签相关
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags", tag.List)

		// 文章相关
		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles", article.List)
	}

	return r
}
