package swagger

type ForbiddenResponse struct {
	Message string `json:"message" example:"You don't have permission to access this resource"`
	Code    string `json:"code" example:"FORBIDDEN"`
}
