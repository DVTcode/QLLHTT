package routes

import (
	"QLLHTT/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Group API
	api := r.Group("/api")

	// Auth routes (cho cả sinh viên và giảng viên)
	api.POST("/register", handlers.Register)
	api.POST("/login", handlers.Login)

	// Student routes
	student := api.Group("/student")
	student.Use(handlers.AuthMiddleware("student"))
	{
		student.GET("/courses", handlers.GetAllCourses)
		student.GET("/materials/:courseId", handlers.GetCourseMaterials)
		student.GET("/enrollments", handlers.GetEnrollmentsForStudent) // Lấy danh sách các khóa học đã đăng ký
	}

	// Teacher routes
	teacher := api.Group("/teacher")
	teacher.Use(handlers.AuthMiddleware("teacher"))
	{
		teacher.POST("/courses", handlers.CreateCourse)
		teacher.PUT("/courses/:id", handlers.UpdateCourse)
		teacher.DELETE("/courses/:id", handlers.DeleteCourse)
		teacher.GET("/courses", handlers.FilterCourses)

		teacher.POST("/materials", handlers.UploadMaterial)
		teacher.PUT("/materials/:id", handlers.UpdateMaterial)
		teacher.DELETE("/materials/:id", handlers.DeleteMaterial)
	}

	// Enrollment routes
	api.POST("/enrollments", handlers.EnrollInCourse)                                   // Sinh viên đăng ký khóa học
	api.DELETE("/enrollments/:id", handlers.UnenrollFromCourse)                         // Sinh viên hủy đăng ký
	api.GET("/teacher/courses/:courseID/enrollments", handlers.GetEnrollmentsForCourse) // Lấy danh sách sinh viên đăng ký khóa học
}
