package response

type UserResponse struct {
	Email string `json:"email"`
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int8   `json:"age"`
}
