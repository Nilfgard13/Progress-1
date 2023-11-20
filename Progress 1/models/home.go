package models

type Home struct {
	Id         int64  `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"type:varchar(30)" json:"name"`
	Deadline   string `gorm:"type:date" json:"deadline"`
	Priority   string `gorm:"type:varchar(30)" json:"priority"`
	Created_At string `gorm:"type:date" json:"created_at"`
}
