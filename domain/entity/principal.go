package entity

type Principal struct {
	Id       int `gorm:"primary_key"json:"id"`
	UserId   int `gorm:"index:user_id"json:"user_id"`
	GroupId int `gorm:"index:domain_id"json:"group_id"`
}

func (m Principal) TableName() string {
	return "principals"
}
