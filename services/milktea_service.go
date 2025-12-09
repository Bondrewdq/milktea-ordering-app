package services

import (
	"context"
	
	"github.com/Bondrewdq/milktea-ordering-app/models"
	"gorm.io/gorm"
)

type MilkTeaService interface {
	GetAllMilkTeas(ctx context.Context) ([]models.Milktea, error)
	GetMilkTeaByID(ctx context.Context, id uint) (*models.Milktea, error)
	CreateMilkTea(ctx context.Context, milktea *models.Milktea) error
	UpdateMilkTea(ctx context.Context, id uint, milktea *models.Milktea) error
	DeleteMilkTea(ctx context.Context, id uint) error
}

type milkTeaService struct {
	db *gorm.DB
}

func NewMilkTeaService(db *gorm.DB) MilkTeaService {
	return &milkTeaService{db: db}
}

func (s *milkTeaService) GetAllMilkTeas(ctx context.Context) ([]models.Milktea, error) {
	var milkteas []models.Milktea
	if err := s.db.WithContext(ctx).Find(&milkteas).Error; err != nil {
		return nil, err
	}
	return milkteas, nil
}

func (s *milkTeaService) GetMilkTeaByID(ctx context.Context, id uint) (*models.Milktea, error) {
	var milktea models.Milktea
	if err := s.db.WithContext(ctx).First(&milktea, id).Error; err != nil {
		return nil, err
	}
	return &milktea, nil
}

func (s *milkTeaService) CreateMilkTea(ctx context.Context, milktea *models.Milktea) error {
	return s.db.WithContext(ctx).Create(milktea).Error
}

func (s *milkTeaService) UpdateMilkTea(ctx context.Context, id uint, milktea *models.Milktea) error {
	return s.db.WithContext(ctx).Model(&models.Milktea{}).Where("id = ?", id).Updates(milktea).Error
}

func (s *milkTeaService) DeleteMilkTea(ctx context.Context, id uint) error {
	return s.db.WithContext(ctx).Delete(&models.Milktea{}, id).Error
}