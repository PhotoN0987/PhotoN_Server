package model

import "time"

// CommonColumn すべてのDBに共通する項目
type CommonColumn struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreatedResponse 作成時の返却値
type CreatedResponse struct {
	ID int64 `json:"id"`
}
