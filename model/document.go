package model

import "knit-integration/model/enum"

type FileData struct {
	ID       string        `json:"id"`
	Category string        `json:"category"`
	FileName string        `json:"fileName"`
	Comment  string        `json:"comment"`
	Content  string        `json:"content"`
	Encoding enum.Encoding `json:"encoding"`
	URL      string        `json:"url"`
}
