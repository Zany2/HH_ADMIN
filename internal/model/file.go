package model

type FileUploadRes struct {
	Id       int64  `json:"id"`
	FileName string `json:"file_name"`
	FileUrl  string `json:"file_url"`
}
