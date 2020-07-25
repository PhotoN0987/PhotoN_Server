package model

// User ユーザー情報
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Introduce string `json:"introduce"`
	Password  string `json:"password"`
}

// UserTable users table type
type UserTable struct {
	User
	CommonColumn
}
