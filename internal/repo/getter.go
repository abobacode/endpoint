package repo

import (
	"context"
	"encoding/json"
	"github.com/abobacode/endpoint/pkg/database"
	"time"

	builder "github.com/doug-martin/goqu/v9"
)

type Getter struct {
	database database.Pool
}

func (g *Getter) GetDateFromDataBase(ctx context.Context, id int) (time.Time, error) {
	var date time.Time

	_, err := g.database.Builder().
		Select(builder.L("updated_at")).
		Where(builder.C("project_id").Eq(id)).
		From("video").
		ScanValContext(ctx, &date)
	if err != nil {
		return time.Time{}, err
	}

	return date, nil
}

func (g *Getter) GetCategoriesFromDataBase(ctx context.Context, id int) ([]string, error) {
	var categories []string
	var categoriesJSON string

	_, err := g.database.Builder().
		Select(builder.L("categories")).
		Where(builder.C("project_id").Eq(id)).
		From("video").
		ScanValContext(ctx, &categoriesJSON)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal([]byte(categoriesJSON), &categories); err != nil {
		return nil, err
	}

	return categories, nil
}

func NewGetter(db database.Pool) *Getter {
	return &Getter{database: db}
}
