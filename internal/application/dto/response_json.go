package dto

type SuccessResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error bool   `json:"error"`
	Msg   string `json:"msg"`
}
