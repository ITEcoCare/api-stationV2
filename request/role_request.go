package request

type RequestIdRole struct {
	ID int `uri:"id" binding:"required,number"`
}
