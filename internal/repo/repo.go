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

func (r *Repo) FetchAdsBlock(ctx context.Context, id int) ([]models.AdsStruct, error) {
	var ads []models.AdsStruct

	err := r.database.Builder().
		Select(builder.L("*")).
		From("ads").
		Where(builder.C("video_id").Eq(id)).
		ScanStructsContext(ctx, &ads)
	if err != nil {
		return nil, err
	}

	return ads, nil
}

func New(db database.Pool) *Repo {
	return &Repo{
		database: db,
	}
}
