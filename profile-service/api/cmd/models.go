package cmd

type Profile struct {
	Avatar   string `json:"avatar"`
	Name     string `json:"name"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Cpf      int64  `json:"cpf"`
	Number   int64  `json:"number"`
	Password string `json:"password"`
}

type Job struct {
	Id          uint64  `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Number      uint64  `json:"number"`
	Category    string  `json:"category"`
	Rate        float64 `json:"rate"`
}
