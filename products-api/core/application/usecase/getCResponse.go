package usecase

import "github.com/products-api/core/domain/models"

type ICResponse interface {
	BuildResponse(msg string, status string) models.CResponse
}

type CResponse struct{}

func BuildResponse(msg string, status int) models.CResponse {
	return models.CResponse{
		Message: msg,
		Status:  status,
	}
}
