package models

import (
	"gorm.io/gorm"
)

// Mô hình Course
type Course struct {
	gorm.Model
	Name        string `json:"name"`        // Tên khóa học
	Description string `json:"description"` // Mô tả khóa học
	TeacherID   uint   `json:"teacher_id"`  // ID của giảng viên (trong User)

	// Mối quan hệ với User (Giảng viên)
	Teacher User `gorm:"foreignKey:TeacherID"`

	// Mối quan hệ với Enrollment
	Enrollments []Enrollment `json:"enrollments" gorm:"foreignKey:CourseID"` // Danh sách sinh viên đã đăng ký khóa học
}
