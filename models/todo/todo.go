package todo

var Todos []Todo

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
