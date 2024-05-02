package repo

import (
	"context"

	builder "github.com/doug-martin/goqu/v9"

	"github.com/abobacode/endpoint/internal/models"
	"github.com/abobacode/endpoint/pkg/database"
)

type Repo struct {
	database database.Pool
}

func (r *Repo) FetchDataBlock(ctx context.Context, id int) ([]models.DataStruct, error) {
	var data []models.DataStruct

	err := r.database.Builder().
		Select(builder.L("*")).
		From("data").
		Where(builder.C("id").Eq(id)).
		ScanStructsContext(ctx, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (r *Repo) SaveDataBlock(ctx context.Context, data []models.DataStruct) error {
	_, err := r.database.Builder().
		Insert("youtube").
		Rows(data).
		Executor().
		ExecContext(ctx)
	if err != nil {
		return err
	}

	return err
}

func New(db database.Pool) *Repo {
	return &Repo{
		database: db,
	}
}
