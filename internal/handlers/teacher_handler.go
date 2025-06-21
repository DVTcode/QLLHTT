package handlers

import (
	"net/http"

	"QLLHTT/internal/config"
	"QLLHTT/internal/models"
	"QLLHTT/internal/utils"

	"github.com/gin-gonic/gin"
)

// ========================== TEACHER ==========================

func CreateCourse(c *gin.Context) {
	user, _ := utils.GetCurrentUser(c)
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	input.TeacherID = user.ID
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create course"})
		return
	}
	c.JSON(http.StatusOK, input)
}

func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var input models.Course
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := config.DB.Model(&models.Course{}).Where("id = ?", id).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	var course models.Course

	// 🔒 Lấy user_id từ token
	teacherID, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	// 🔍 Tìm course theo ID và teacherID
	if err := config.DB.Where("id = ? AND teacher_id = ?", id, teacherID).First(&course).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Course not found or unauthorized"})
		return
	}

	// ❌ Soft delete (nếu dùng gorm.Model)
	if err := config.DB.Delete(&course).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}

// ✅ Filter courses based on query parameters
func FilterCourses(c *gin.Context) {
	// Lấy tham số từ query parameters
	title := c.DefaultQuery("title", "")          // Lọc theo tiêu đề khóa học
	teacherID := c.DefaultQuery("teacher_id", "") // Lọc theo ID của giảng viên (nếu có)

	var courses []models.Course

	// Lọc theo tiêu chí
	query := config.DB.Model(&models.Course{})

	// Nếu có tiêu đề, lọc theo tiêu đề
	if title != "" {
		query = query.Where("title LIKE ?", "%"+title+"%")
	}

	// Nếu có teacher_id, lọc theo giảng viên
	if teacherID != "" {
		query = query.Where("teacher_id = ?", teacherID)
	}

	// Thực hiện truy vấn và trả về kết quả
	if err := query.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}

	// Trả về danh sách các khóa học
	c.JSON(http.StatusOK, gin.H{"courses": courses})
}

func UploadMaterial(c *gin.Context) {
	var input models.Document
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := config.DB.Create(&input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to upload"})
		return
	}
	c.JSON(http.StatusOK, input)
}

func UpdateMaterial(c *gin.Context) {
	id := c.Param("id")
	var input models.Document
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	if err := config.DB.Model(&models.Document{}).Where("id = ?", id).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update material"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "material updated"})
}

func DeleteMaterial(c *gin.Context) {
	id := c.Param("id")
	if err := config.DB.Delete(&models.Document{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete material"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "material deleted"})
}
