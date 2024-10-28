package handler

import "github.com/gin-gonic/gin"

type Operation interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}
