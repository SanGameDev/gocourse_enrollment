package enrollment

import (
	"context"
	"log"

	"github.com/SanGameDev/gocourse_domain/domain"
)

type (
	Filters struct {
		UserID   string
		CourseID string
	}

	Service interface {
		Create(ctx context.Context, userID, courseID string) (*domain.Enrollment, error)
	}
	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(log *log.Logger, repo Repository) Service {
	return &service{
		log:  log,
		repo: repo,
	}
}

func (s *service) Create(ctx context.Context, userID, courseID string) (*domain.Enrollment, error) {

	enroll := &domain.Enrollment{
		UserID:   userID,
		CourseID: courseID,
		Status:   "pending",
	}

	if err := s.repo.Create(ctx, enroll); err != nil {
		s.log.Println(err)
		return nil, err
	}

	return enroll, nil
}
