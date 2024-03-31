package repository

import (
	"construction-organization-system/internal/pkg/model"
	"context"
)

type EngineerWorkerRepository interface {
	Save(ctx context.Context, entity model.EngineerWorker) (int, error)
	Find(ctx context.Context, id int) (*model.EngineerWorker, error)
	Update(ctx context.Context, entity model.EngineerWorker) error
	Delete(ctx context.Context, id int) error
}
