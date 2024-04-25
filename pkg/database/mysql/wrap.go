package mysql

import (
	"context"
	"database/sql"
	"time"

	builder "github.com/doug-martin/goqu/v9"
	// nolint:golint // it's OK
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/abobacode/endpoint/pkg/database"
	"github.com/abobacode/endpoint/pkg/log"
)

type ConnectionPool struct {
	db *builder.Database
}

func (c *ConnectionPool) Builder() *builder.Database {
	return c.db
}

// Drop close not implemented in database
func (c *ConnectionPool) Drop() error {
	return nil
}

func (c *ConnectionPool) DropMsg() string {
	return "close database: this is not implemented"
}

func New(ctx context.Context, opt *database.Opt) (*ConnectionPool, error) {
	opt.UnwrapOrPanic()

	db, err := sql.Open(opt.Dialect, opt.ConnectionString())
	if err != nil {
		return nil, err
	}

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err = db.PingContext(pingCtx); err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(opt.MaxIdleConns)
	db.SetMaxOpenConns(opt.MaxOpenConns)
	db.SetConnMaxLifetime(opt.MaxConnMaxLifetime)

	dialect := builder.Dialect(opt.Dialect)
	connect := &ConnectionPool{
		db: dialect.DB(db),
	}

	if opt.Debug {
		logger := &database.Logger{}
		logger.SetCallback(func(format string, v ...interface{}) {
			log.Info(v)
		})
		connect.db.Logger(logger)
	}

	return connect, nil
}
