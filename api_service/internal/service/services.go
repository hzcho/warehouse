package service

import (
	"api_service/internal/config"
	"api_service/internal/domain/service"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Services struct {
	service.Auth
	service.Product
	service.Audit
	service.Category
}

func NewServices(urls config.URL, client *http.Client, log *logrus.Logger) *Services {
	return &Services{
		Auth:     NewAuth(urls.Auth, client, log),
		Product:  NewProduct(urls.Warehouse, client, log),
		Audit:    NewAudit(urls.Audit, client, log),
		Category: NewCategory(urls.Warehouse, client, log),
	}
}
