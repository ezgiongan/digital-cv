package db

import "digital-cv-backend/internal/models"

type CVRepository interface {
	GetCV() (*models.CV, error)
}