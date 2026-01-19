package model

type Article struct {
	Id         string  `db:"id" json:"id"`
	Num        int64   `db:"num" json:"num"`
	Title      string  `db:"title" json:"title"`
	Writer     string  `db:"writer" json:"writer"`
	WriteDate  string  `db:"write_date" json:"write_date" swaggertype:"string"`
	ArticleUrl string  `db:"article_url" json:"article_url"`
	Content    string  `db:"content" json:"content"`
	IsNotice   bool    `db:"is_notice" json:"is_notice"`
	Files      []Files `json:"files"`
}
