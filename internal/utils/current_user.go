// File: internal/utils/current_user.go
package utils

import (
	"QLLHTT/internal/config"
	"QLLHTT/internal/models"
	"errors"

	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (models.User, error) {
	val, exists := c.Get("user_id")
	if !exists {
		return models.User{}, errors.New("user_id not found in context")
	}

	var user models.User
	if err := config.DB.First(&user, val).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}
