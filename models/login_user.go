package models

type LoginUser struct {
	Username   string // 中文名
	EmployeeId string // 工号
	EnUsername string // 英文名
	Email      string // 邮箱
	IsAdmin    bool
}
