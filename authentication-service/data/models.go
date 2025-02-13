package data

type User struct {
	Id       uint64 `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Cpf      int64  `json:"cpf"`
	Number   int64  `json:"number"`
	Password string `json:"password"`
}
