package payload

type GetUser struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
}

type GetTask struct {
	ID           uint   `json:"id" form:"id"`
	Title        string `json:"title" form:"title"`
	Description  string `json:"description" form:"description"`
	Status       string `json:"status" form:"status"`
	Category     string `json:"category" form:"category"`
	Owner        string `json:"owner" form:"owner"`
	Collabolator []Collaborator
}

type Collaborator struct {
	ID       uint   `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
}

type User struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
