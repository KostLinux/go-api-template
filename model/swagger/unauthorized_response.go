package swagger

type UnauthorizedResponse struct {
	Message string `json:"message" example:"You are not authorized to access this resource"`
	Code    string `json:"code" example:"UNAUTHORIZED"`
}
