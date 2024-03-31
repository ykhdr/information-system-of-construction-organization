package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type EngineerTeamRepository interface {
	Save(ctx context.Context, entity model.EngineerTeam) (int, error)
	Find(ctx context.Context, id int) (*model.EngineerTeam, error)
	Update(ctx context.Context, entity model.EngineerTeam) error
	Delete(ctx context.Context, id int) error
}
