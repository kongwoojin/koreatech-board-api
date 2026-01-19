package model

type Files struct {
	FileName string `db:"file_name" json:"file_name"`
	FileUrl  string `db:"file_url" json:"file_url"`
}
