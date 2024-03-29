package model

import "github.com/edgedb/edgedb-go"

type Board struct {
	Id        edgedb.UUID      `edgedb:"id" json:"id"`
	Num       string           `edgedb:"num" json:"num"`
	Title     string           `edgedb:"title" json:"title"`
	Writer    string           `edgedb:"writer" json:"writer"`
	WriteDate edgedb.LocalDate `edgedb:"write_date" json:"write_date"`
	ReadCount int64            `edgedb:"read_count" json:"read_count"`
}
