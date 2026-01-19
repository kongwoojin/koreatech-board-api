package model

type ApiArticle struct {
	StatusCode int     `json:"status_code"` // Status code of request
	Error      string  `json:"error"`       // Error message
	Num        int64   `json:"num"`
	Id         string  `json:"id"`
	Title      string  `json:"title"`
	Writer     string  `json:"writer"`
	WriteDate  string  `json:"write_date" swaggertype:"string"`
	ArticleUrl string  `json:"article_url"`
	Content    string  `json:"content"`
	IsNotice   bool    `json:"is_notice"`
	Files      []Files `json:"files"`
}
