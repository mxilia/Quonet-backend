package dto

type CreateUserByGoogleRequest struct {
	ProfileUrl string `json:"profile_url,omitempty"`
	Email      string `json:"email" validate:"required"`
}

type UserPatchRequest struct {
	Handler    string `json:"handler,omitempty"`
	ProfileUrl string `json:"profile_url,omitempty"`
	Role       string `json:"role,omitempty"`
}
