package enrollment

import (
	"context"
	"log"

	"github.com/SanGameDev/gocourse_domain/domain"
	"gorm.io/gorm"
)

type (
	Repository interface {
		Create(ctx context.Context, enroll *domain.Enrollment) error
	}

	repo struct {
		db  *gorm.DB
		log *log.Logger
	}
)

func NewRepo(db *gorm.DB, log *log.Logger) Repository {
	return &repo{
		db:  db,
		log: log,
	}
}

func (repo *repo) Create(ctx context.Context, enroll *domain.Enrollment) error {

	if err := repo.db.WithContext(ctx).Create(enroll).Error; err != nil {
		repo.log.Printf("error: %v", err)
		return err
	}

	repo.log.Println("Enrollment created with id:", enroll.ID)
	return nil
}
