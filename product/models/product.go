/*
 * @Author: lihuan
 * @Date: 2021-12-20 20:29:35
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-21 21:07:38
 * @Email: 17719495105@163.com
 */

package models

type Product struct {
	ID        int64
	DetailId  int64
	CommentId int64
	Url       string
	Des       string
	Name      string
	ShortName string
	Type      int32
	Price     string
	Sale      int64
	Inventory int64
	Score     string
	Discount  string
	CreatedAt int32
	UpdatedAt int32
}

func (p Product) TableName() string {
	return "hc_product"
}
