package model

type Board struct {
	Id        string `db:"id" json:"id"`
	Num       int64  `db:"num" json:"num"`
	Title     string `db:"title" json:"title"`
	Writer    string `db:"writer" json:"writer"`
	WriteDate string `db:"write_date" json:"write_date" swaggertype:"string"`
	ReadCount int64  `db:"read_count" json:"read_count"`
	IsNew     bool   `db:"is_new" json:"is_new"`
	IsNotice  bool   `db:"is_notice" json:"is_notice"`
}
