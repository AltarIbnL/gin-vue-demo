package dto

import "myproject/Model"

type UserDto struct {
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}

//只希望返回用户的姓名和电话
func ToUserDto(user Model.User) UserDto {
	return UserDto{
		Name:      user.Name,
		Telephone: user.Telephone,
	}
}
