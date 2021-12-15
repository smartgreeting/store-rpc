/*
 * @Author: lihuan
 * @Date: 2021-12-14 20:22:15
 * @LastEditors: lihuan
 * @LastEditTime: 2021-12-15 22:51:33
 * @Email: 17719495105@163.com
 */
package models

type User struct {
	ID        uint64
	Username  string
	Password  string
	Avatar    string
	Gender    int32
	Phone     string
	Email     string
	Address   string
	Hobbies   string
	Deleted   int32
	CreatedAt int32
	UpdatedAt int32
}

func (u User) TableName() string {
	return "hc_user"
}
