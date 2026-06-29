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
	err := ValidateCreateApplication(req)
	if err != nil {
		return nil, err
	}

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

func (s *Service) Put(ctx context.Context, id uuid.UUID, req PutApplication) (*Application, error) {
	err := ValidatePutApplication(req)
	if err != nil {
		return nil, err
	}

	app, err := s.r.Put(ctx, id, req)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (s *Service) Patch(ctx context.Context, id uuid.UUID, req PatchApplication) (*Application, error) {
	app, err := s.r.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		app.Name = *req.Name
	}

	if req.Namespace != nil {
		app.Namespace = *req.Namespace
	}

	if req.Image != nil {
		app.Image = *req.Image
	}

	if req.Replicas != nil {
		app.Replicas = *req.Replicas
	}

	if req.Port != nil {
		app.Port = *req.Port
	}

	app, err = s.r.Patch(ctx, app)
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
