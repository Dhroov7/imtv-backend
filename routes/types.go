package routes

type login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
