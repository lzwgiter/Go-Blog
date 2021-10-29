// @Author: float311@163.com
// @Description: 文章标签关联对象
// @Version: 1.0.0
// Date: 2021/10/25 21:26
// Update: 2021/10/25 21:26

package model

type ArticleTag struct {
	*Model
	ArticleId uint32 `json:"article_id"`
	TagId     uint32 `json:"tag_id"`
}

func (a ArticleTag) TableName() string {
	return "blog_article_tag"
}
