package response

type RepositoryResult struct {
	Result interface{}
	Error  error
}

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
