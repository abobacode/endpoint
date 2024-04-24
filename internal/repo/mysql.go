package repo

import (
	"context"
	"github.com/abobacode/endpoint/internal/models"
	"github.com/abobacode/endpoint/pkg/database"

	builder "github.com/doug-martin/goqu/v9"
)

type Repository struct {
	driver database.Driver
}

func (r *Repository) FetchEpg(ctx context.Context) ([]models.Epg, error) {
	var epg []models.Epg

	err := r.driver.Conn().
		Select(builder.L("*")).
		From("epg").
		Where(builder.C("video_url").Neq("")).
		ScanStructsContext(ctx, &epg)
	if err != nil {
		return nil, err
	}

	return epg, nil
}

func NewMySQL(driver database.Driver) *Repository {
	return &Repository{
		driver: driver,
	}
}
