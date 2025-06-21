package handlers

import (
	"QLLHTT/internal/config"
	"QLLHTT/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// POST /enrollments
func EnrollInCourse(c *gin.Context) {
	var input struct {
		StudentID uint `json:"student_id" binding:"required"`
		CourseID  uint `json:"course_id" binding:"required"`
	}

	// Bind input data from JSON
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Kiểm tra xem khóa học có tồn tại không
	var course models.Course
	if err := config.DB.Where("id = ?", input.CourseID).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
		return
	}

	// Kiểm tra xem sinh viên có tồn tại không
	var student models.User
	if err := config.DB.Where("id = ?", input.StudentID).First(&student).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Student not found"})
		return
	}

	// Kiểm tra xem sinh viên đã đăng ký khóa học này chưa
	var existingEnrollment models.Enrollment
	if err := config.DB.Where("student_id = ? AND course_id = ?", input.StudentID, input.CourseID).First(&existingEnrollment).Error; err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Student already enrolled in this course"})
		return
	}

	// Đăng ký sinh viên vào khóa học
	enrollment := models.Enrollment{
		StudentID: input.StudentID,
		CourseID:  input.CourseID,
		JoinedAt:  time.Now(),
	}

	if err := config.DB.Create(&enrollment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to enroll student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Enrollment successful", "enrollment": enrollment})
}

// DELETE /enrollments/:id
func UnenrollFromCourse(c *gin.Context) {
	id := c.Param("id")

	// Kiểm tra xem bản ghi có tồn tại không
	var enrollment models.Enrollment
	if err := config.DB.Where("id = ?", id).First(&enrollment).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Enrollment not found"})
		return
	}

	// Hủy đăng ký sinh viên khỏi khóa học
	if err := config.DB.Delete(&models.Enrollment{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete enrollment"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully unenrolled"})
}

// GET /student/enrollments
func GetEnrollmentsForStudent(c *gin.Context) {
	studentID := c.MustGet("user_id").(uint) // Lấy ID của sinh viên từ token (JWT)
	var enrollments []models.Enrollment

	// Truy vấn danh sách các khóa học đã đăng ký cho sinh viên
	if err := config.DB.Where("student_id = ?", studentID).Preload("Course").Preload("Student").Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	// Trả về danh sách các khóa học đã đăng ký
	c.JSON(http.StatusOK, gin.H{"enrollments": enrollments})
}

// GET /teacher/courses/:courseID/enrollments
func GetEnrollmentsForCourse(c *gin.Context) {
	courseID := c.Param("courseID") // Lấy ID khóa học từ URL
	var enrollments []models.Enrollment

	// Lấy danh sách sinh viên đã đăng ký khóa học này
	if err := config.DB.Where("course_id = ?", courseID).Preload("Student").Find(&enrollments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch enrollments"})
		return
	}

	// Trả về danh sách sinh viên đã đăng ký khóa học
	c.JSON(http.StatusOK, gin.H{"enrollments": enrollments})
}
