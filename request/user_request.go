package request

type RequestId struct {
	ID int `uri:"id" binding:"required,number"`
}

type RequestLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CheckEmailRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type CheckUsernameRequest struct {
	Username string `json:"username" binding:"required"`
}
