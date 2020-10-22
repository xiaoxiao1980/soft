package model

type Filepath struct {
	Path string `json:"path" binding:"required"`
}
