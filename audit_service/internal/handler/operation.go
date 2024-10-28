package handler

import (
	"audit/internal/domain/net/request"
	"audit/internal/domain/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation struct {
	operUseCase usecase.Operation
}

func NewOperations(operUseCase usecase.Operation) *Operation {
	return &Operation{
		operUseCase: operUseCase,
	}
}

func (h *Operation) GetById(c *gin.Context) {
	stringID := c.Param("id")
	if stringID == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, "empty id")
		return
	}

	id, err := primitive.ObjectIDFromHex(stringID)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "incorrect id")
		return
	}

	operation, err := h.operUseCase.GetById(c, id)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, operation)
}

func (h *Operation) GetAll(c *gin.Context) {
	var filter request.GetAllFilter

	if err := c.ShouldBind(filter); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, "bad request")
	}

	operation, err := h.operUseCase.GetAll(c, filter)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, operation)
}
