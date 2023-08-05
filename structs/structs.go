package structs

type User struct {
	UUID     string `json:"uuid"`
	Nickname string `json:"nickname"`
	Password string `json:"password"`
}

type UserResponse struct {
	UUID     string `json:"uuid"`
	Nickname string `json:"nickname"`
}

type Response struct {
	Status string        `json:"status"`
	Body   *UserResponse `json:"body,omitempty"`
	Error  string        `json:"error,omitempty"`
}
