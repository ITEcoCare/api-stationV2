package request

type RequestIdModuleApp struct {
	ID int `uri:"id" binding:"required,number"`
}
