package request

type GetAuthorRequest struct {
	ID int `query:"id"`
}

type CreateAuthorRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
