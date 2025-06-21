package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(100)" json:"username"`     // tốt hơn longtext
	Email    string `gorm:"type:varchar(191);unique" json:"email"` // đảm bảo đồng nhất
	Password string `gorm:"type:varchar(255)" json:"-"`            // không trả ra client
	Role     string `gorm:"type:varchar(50)" json:"role"`          // chỉ student/teacher
}
