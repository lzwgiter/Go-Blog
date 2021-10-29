// @Author: float311@163.com
// @Description: 文章对象
// @Version: 1.0.0
// Date: 2021/10/25 21:22
// Update: 2021/10/25 21:22

package model

type Article struct {
	*Model
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CoverImageURL string `json:"cover_image_url"`
	State         uint8  `json:"state"`
}

func (a Article) TableName() string {
	return "blog_article"
}
