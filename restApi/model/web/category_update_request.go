package web

type CategoryUpdateRequest struct {
	Id   int    `validate:"required" json:"Id"`
	Name string `validate:"required" json:"Name"`
}
