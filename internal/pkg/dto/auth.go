package dto

type LoginInfo struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JWT struct {
	Token string `json:"token"`
}

type DBUser struct {
	ID           uint64
	Login        string
	PasswordHash string `db:"password_hash"`
	IsAdmin      bool   `db:"is_admin"`
}
