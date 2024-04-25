package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strings"
	"time"

	builder "github.com/doug-martin/goqu/v9"
	// nolint:golint // reason
	_ "github.com/doug-martin/goqu/v9/dialect/mysql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/abobacode/endpoint/pkg/xerror"
)

type Driver interface {
	Conn() *builder.Database
}

type Connector struct {
	db      *builder.Database
	rawConn *sql.DB
}

func (c *Connector) Conn() *builder.Database {
	return c.db
}

// Drop close not implemented in database
func (c *Connector) Drop() error {
	return c.rawConn.Close()
}

func (c *Connector) DropMsg() string {
	return "close mysql connection"
}

func NewMySQLDriver(ctx context.Context, cfg *DriverConfiguration) (*Connector, error) {
	err := cfg.checkBeforeInitializing()
	if err != nil {
		return nil, err
	}
	db, err := sql.Open(cfg.Dialect, buildConnectionString(cfg))
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetConnMaxLifetime(cfg.MaxConnMaxLifetime)
	dialect := builder.Dialect(cfg.Dialect)
	connection := &Connector{
		db:      dialect.DB(db),
		rawConn: db,
	}
	if cfg.Debug {
		dbLogger := &Logger{}
		dbLogger.SetCallback(func(format string, v ...interface{}) {
			log.Println(v)
		})
		connection.db.Logger(dbLogger)
	}
	return connection, nil
}

type DriverConfiguration struct {
	Host               string        `yaml:"host"`
	User               string        `yaml:"user"`
	Password           string        `yaml:"password"`
	Port               string        `yaml:"port"`
	Name               string        `yaml:"name"`
	Dialect            string        `yaml:"dialect"`
	Debug              bool          `yaml:"debug"`
	MaxIdleConns       int           `yaml:"max_idle_conns"`
	MaxOpenConns       int           `yaml:"max_open_conns"`
	MaxConnMaxLifetime time.Duration `yaml:"max_conn_max_lifetime"`
	Collation          string        `yaml:"collation"`
}

func (o *DriverConfiguration) checkBeforeInitializing() error {
	if o.Dialect == "" {
		o.Dialect = "mysql"
	}
	if o.Host == "" {
		o.Host = "@"
	}
	if strings.EqualFold(o.Host, "") {
		o.Host = "@"
	} else if !strings.Contains(o.Host, "@") {
		o.Host = fmt.Sprintf("@tcp(%s)", o.Host)
	}

	if o.MaxIdleConns <= 0 {
		return xerror.ErrMaxIdleConns
	}
	if o.MaxOpenConns <= 0 {
		return xerror.ErrMaxOpenConns
	}
	if o.MaxConnMaxLifetime <= 0 {
		return xerror.ErrMaxConnMaxLifetime
	}
	return nil
}

func buildConnectionString(cfg *DriverConfiguration) string {
	return fmt.Sprintf("%s:%s%s/%s?parseTime=true",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Name,
	)
}

type MockDriver struct{}

func (m *MockDriver) Conn() *builder.Database {
	return nil
}

func (m *MockDriver) Drop() error {
	return nil
}

func (m *MockDriver) DropMsg() string {
	return "close mock mysql connection"
}
