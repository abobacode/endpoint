package service

import (
	"context"

	"github.com/abobacode/endpoint/pkg/database"
	"github.com/abobacode/endpoint/pkg/database/mysql"
	"github.com/abobacode/endpoint/pkg/drop"
)

type Options struct {
	Database *database.Opt
}

type Endpoint struct {
	*drop.Impl
	Pool database.Pool
}

func New(ctx context.Context, opt *Options) (*Endpoint, error) {
	var err error
	endpoint := &Endpoint{}
	endpoint.Impl = drop.NewContext(ctx)

	if opt.Database != nil {
		endpoint.Pool, err = mysql.New(endpoint.Context(), opt.Database)
		if err != nil {
			return nil, err
		}
		endpoint.AddDropper(endpoint.Pool.(*mysql.ConnectionPool))
	}

	return endpoint, nil
}
