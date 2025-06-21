package models

import (
	"time"

	"gorm.io/gorm"
)

// Mô hình Enrollment
type Enrollment struct {
	gorm.Model
	StudentID uint      `gorm:"not null" json:"student_id"` // ID của sinh viên
	CourseID  uint      `gorm:"not null" json:"course_id"`  // ID của khóa học
	JoinedAt  time.Time `gorm:"not null" json:"joined_at"`  // Thời gian đăng ký

	// Mối quan hệ với Course
	Course Course `json:"course" gorm:"foreignKey:CourseID"`

	// Mối quan hệ với User (Sinh viên)
	Student User `json:"student" gorm:"foreignKey:StudentID"`
}
