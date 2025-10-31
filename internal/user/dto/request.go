package dto

type CreateUserByGoogleRequest struct {
	ProfileUrl string `json:"picture"`
	Email      string `json:"email" validate:"required"`
}

type UserPatchRequest struct {
	Handler    string `json:"handler"`
	ProfileUrl string `json:"profile_url"`
}
