package model

// User はユーザードメインモデルです。
type User struct {
	ID        string
	Name      string
	Email     string
	CreatedAt string
	UpdatedAt string
}

// CreateUserInput はユーザー作成時の入力データです。
type CreateUserInput struct {
	Name  string
	Email string
}
