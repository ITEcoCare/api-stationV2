package response

type ValidationResponse struct {
	Success     bool         `json:"success"`
	Validations []Validation `json:"validations"`
}

type Validation struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
