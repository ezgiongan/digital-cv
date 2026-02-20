package services
import (
	"digital-cv-backend/internal/db"
	"digital-cv-backend/internal/models"
)

type CVService struct {
	repo db.CVRepository
}

func NewCVService(repo *db.PostgresCVRepository) *CVService {
	return &CVService{repo: repo}
}

func (s *CVService) GetCV() (*models.CV, error) {
	return s.repo.GetCV()
}