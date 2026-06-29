package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/hugaojanuario/devplatform/internal/errs"
)

type Service struct {
	r ApplicationRepository
}

func NewService(r ApplicationRepository) *Service {
	return &Service{r: r}
}

func (s *Service) Create(ctx context.Context, req CreateApplication) (*Application, error) {

	app, err := s.r.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *Service) FindAll(ctx context.Context) ([]Application, error) {
	apps, err := s.r.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	return apps, nil
}

func (s *Service) FindByID(ctx context.Context, id uuid.UUID) (*Application, error) {
	app, err := s.r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if app == nil {
		return nil, errs.NotFound("Application")
	}

	return app, nil
}

func (s *Service) Update(ctx context.Context, id uuid.UUID, req UpdateApplication) (*Application, error) {
	app, err := s.r.Update(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *Service) Delete(ctx context.Context, id uuid.UUID) error {
	err := s.r.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
