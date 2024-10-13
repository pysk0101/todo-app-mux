package domain

type Todo struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	IsDone     bool   `json:"is_done"`
	Created_At string `json:"created_at"`
}
