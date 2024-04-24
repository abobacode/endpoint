package database

import builder "github.com/doug-martin/goqu/v9"

type Pool interface {
	Builder() *builder.Database
}
