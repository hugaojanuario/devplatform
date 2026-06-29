package application

import (
	"strings"

	"github.com/hugaojanuario/devplatform/internal/errs"
)

func ValidateCreateApplication(req CreateApplication) error {
	if strings.TrimSpace(req.Name) == "" {
		return errs.Required("name")
	}

	if strings.TrimSpace(req.Namespace) == "" {
		return errs.Required("namespace")
	}

	if strings.TrimSpace(req.Image) == "" {
		return errs.Required("image")
	}

	if req.Replicas <= 0 {
		return errs.Invalid("replicas")
	}

	if req.Port <= 0 || req.Port > 65535 {
		return errs.Invalid("port")
	}

	return nil
}

func ValidatePutApplication(req PutApplication) error {
	if strings.TrimSpace(req.Name) == "" {
		return errs.Required("name")
	}

	if strings.TrimSpace(req.Namespace) == "" {
		return errs.Required("namespace")
	}

	if strings.TrimSpace(req.Image) == "" {
		return errs.Required("image")
	}

	if req.Replicas <= 0 {
		return errs.Invalid("replicas")
	}

	if req.Port <= 0 || req.Port > 65535 {
		return errs.Invalid("port")
	}

	return nil
}
