package models

import "gorm.io/gorm"

type Document struct {
	gorm.Model
	Title    string `json:"title"`
	FileURL  string `json:"file_url"`
	Type     string `json:"type"`      // "pdf", "video", "link"
	CourseID uint   `json:"course_id"` // khóa ngoại
}
