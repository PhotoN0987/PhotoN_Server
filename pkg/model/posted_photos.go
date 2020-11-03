package model

// PostedPhotos 投稿写真情報
type PostedPhotos struct {
	PhotoID       int    `json:"photo_id"`
	UserID        string `json:"user_id"`
	PhotoURL      string `json:"photo_url"`
	PhotoComment  string `json:"photo_comment"`
	PhotoCategory string `json:"photo_category"`
	DeleteFlag    string `json:"delete_flag"`
}

// PostedPhotosTable users table type
type PostedPhotosTable struct {
	PostedPhotos
	CommonColumn
}
