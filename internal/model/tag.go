// @Author: float311@163.com
// @Description: 文章标签对象
// @Version: 1.0.0
// Date: 2021/10/25 21:20
// Update: 2021/10/25 21:20

package model

type Tag struct {
	*Model
	Name  string `json:"name"`
	State uint8  `json:"state"`
}

func (t Tag) TableName() string {
	return "blog_tag"
}
