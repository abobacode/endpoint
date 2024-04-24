package database

import (
	"fmt"
	"strings"
	"time"
)

type Opt struct {
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
}

func (o *Opt) UnwrapOrPanic() {
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
		panic("max_idle_conns must be greater than zero")
	}
	if o.MaxOpenConns <= 0 {
		panic("max_open_conns must be greater than zero")
	}
	if o.MaxConnMaxLifetime <= 0 {
		panic("max_conn_max_lifetime must be greater than zero")
	}
}

func (o *Opt) ConnectionString() string {
	return fmt.Sprintf("%s:%s%s/%s?parseTime=true", o.User, o.Password, o.Host, o.Name)
}
