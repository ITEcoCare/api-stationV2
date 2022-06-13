package response

type ValidationResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	// Validations []Validation `json:"validations"`
	Data interface{} `json:"data"`
}

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
