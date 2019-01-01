package entity

type Member struct {
	Id        int       `json:"id"`
	UserId    int       `gorm:"type:varchar(128);not null;index:user_id"json:"user_id"`
	RoleId    int       `gorm:"type:varchar(128);not null;index:role_id"json:"role_id"`
}

func (e Member) TableName() string {
	return "members"
}
