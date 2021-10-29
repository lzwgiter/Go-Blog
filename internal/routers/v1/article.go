// @Author: float311@163.com
// @Description:
// @Version: 1.0.0
// Date: 2021/10/29 12:56
// Update: 2021/10/29 12:56

package v1

import "github.com/gin-gonic/gin"

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (t Article) Get(c *gin.Context)    {}
func (t Article) List(c *gin.Context)   {}
func (t Article) Create(c *gin.Context) {}
func (t Article) Update(c *gin.Context) {}
func (t Article) Delete(c *gin.Context) {}
