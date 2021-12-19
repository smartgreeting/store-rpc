/*
 * @Author: lihuan
 * @Date: 2021-12-19 15:49:48
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-19 16:28:17
 * @Email: 17719495105@163.com
 */
package models

type Banner struct {
	ID        int64
	ProductID int64
	Url       string
	Order     int32
	CreatedAt int32
	UpdatedAt int32
}

func (b Banner) TableName() string {
	return "hc_banner"
}
