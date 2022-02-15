package request

type RequestIdteam struct {
	ID int `uri:"id" binding:"required,number"`
}
