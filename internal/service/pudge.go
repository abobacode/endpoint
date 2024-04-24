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

type Pudge struct {
	*drop.Impl
	driver *database.Connector
	Pool   database.Pool
}

func (p *Pudge) Conn() *database.Connector {
	return p.driver
}

func New(ctx context.Context, opt *Options) (*Pudge, error) {
	var err error
	pudge := &Pudge{}
	pudge.Impl = drop.NewContext(ctx)

	if opt.Database != nil {
		pudge.Pool, err = mysql.New(pudge.Context(), opt.Database)
		if err != nil {
			return nil, err
		}
		pudge.AddDropper(pudge.Pool.(*mysql.ConnectionPool))
	}

	return pudge, nil
}
