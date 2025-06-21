package handlers

import (
	"net/http"

	"QLLHTT/internal/config"
	"QLLHTT/internal/models"

	"github.com/gin-gonic/gin"
)

// ========================== STUDENT ==========================

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	if err := config.DB.Find(&courses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCourseMaterials(c *gin.Context) {
	courseID := c.Param("courseId")
	var documents []models.Document
	if err := config.DB.Where("course_id = ?", courseID).Find(&documents).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch documents"})
		return
	}
	c.JSON(http.StatusOK, documents)
}
