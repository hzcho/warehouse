package handler

import (
	"api_service/internal/domain/net/request"
	"api_service/internal/domain/service"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Audit struct {
	auditService service.Audit
}

func NewAudit(auditService service.Audit) *Audit {
	return &Audit{
		auditService: auditService,
	}
}

func (h *Audit) GetAll(c *gin.Context) {
	token, exists := c.Get("token")
	if !exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empty token"})
		return
	}
	strToken, ok := token.(string)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "token isn't a string"})
		return
	}

	var filter request.GetAllLogsFilter

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
		return
	}

	resp, err := h.auditService.GetAll(c, strToken, filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get logs"})
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	c.Data(resp.StatusCode, resp.Header.Get("Content-Type"), body)
}
