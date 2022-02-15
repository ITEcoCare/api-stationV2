package request

type RequestIdPermission struct {
	ID int `uri:"id" binding:"required,number"`
}
