package aplication

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
)

type ApplicationRepository interface {
	Create(ctx context.Context, req CreateApplication) (*Application, error)
	FindAll(ctx context.Context) ([]Application, error)
	FindByID(ctx context.Context, id uuid.UUID) (*Application, error)
	Update(ctx context.Context, id uuid.UUID, req UpdateApplication) (*Application, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, req CreateApplication) (*Application, error) {
	query := `INSERT INTO applications (name, namespace, image, replicas, port) VALUES ($1, $2, $3, $4, $5) RETURNING id, name, namespace, image, replicas, port, created_at, updated_at`

	application := &Application{}
	err := r.db.QueryRowContext(ctx, query, req.Name, req.Namespace, req.Image, req.Replicas, req.Port).Scan(
		&application.ID,
		&application.Name,
		&application.Namespace,
		&application.Image,
		&application.Replicas,
		&application.Port,
		&application.CreatedAt,
		&application.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (r *repository) FindAll(ctx context.Context) ([]Application, error) {
	query := `SELECT id, name, namespace, image, replicas, port, created_at, updated_at FROM applications`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var applications []Application
	for rows.Next() {
		var application Application
		err := rows.Scan(
			&application.ID,
			&application.Name,
			&application.Namespace,
			&application.Image,
			&application.Replicas,
			&application.Port,
			&application.CreatedAt,
			&application.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		applications = append(applications, application)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return applications, nil
}

func (r *repository) FindByID(ctx context.Context, id uuid.UUID) (*Application, error) {
	query := `SELECT id, name, namespace, image, replicas, port, created_at, updated_at FROM applications WHERE id = $1`
	application := &Application{}
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&application.ID,
		&application.Name,
		&application.Namespace,
		&application.Image,
		&application.Replicas,
		&application.Port,
		&application.CreatedAt,
		&application.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (r *repository) Update(ctx context.Context, id uuid.UUID, req UpdateApplication) (*Application, error) {
	query := `UPDATE applications SET name = $1, namespace = $2, image = $3, replicas = $4, port = $5, updated_at = NOW() WHERE id = $6 RETURNING id, name, namespace, image, replicas, port, created_at, updated_at`
	application := &Application{}
	err := r.db.QueryRowContext(ctx, query, req.Name, req.Namespace, req.Image, req.Replicas, req.Port, id).Scan(
		&application.ID,
		&application.Name,
		&application.Namespace,
		&application.Image,
		&application.Replicas,
		&application.Port,
		&application.CreatedAt,
		&application.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return application, nil
}

func (r *repository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM applications WHERE id = $1`
	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return sql.ErrNoRows // não encontrado
	}

	return nil
}
