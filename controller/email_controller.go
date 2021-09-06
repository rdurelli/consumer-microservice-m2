package controller

import (
	"consumer-rabbitmq/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

type EmailController struct {
	repo repository.Repository
}

func NewEmailController(repo repository.Repository) EmailController {
	return EmailController{
		repo: repo,
	}
}

func (eC EmailController) Find(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	emails, err := eC.repo.Find(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"success": false,
		})
		return
	}
	if len(*emails) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "All email has been successfully sent.",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"emails_not_sent": *emails,
		"size":            len(*emails),
		"success":         true,
	})
}

func (eC EmailController) ShowAll(c *gin.Context) {
	offset := c.DefaultQuery("offset", "0")
	limit := c.DefaultQuery("limit", "10")
	emails, err := eC.repo.ShowAll(offset, limit)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   err,
			"success": false,
		})
		return
	}
	if len(*emails) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "There is none email information in the database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"emails":  *emails,
		"size":    len(*emails),
		"success": true,
	})
}
