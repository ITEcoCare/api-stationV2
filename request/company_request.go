package request

type RequestIdCompany struct {
	ID int `json:"id" binding:"required,number"`
}
