package internal

type UserPublicData struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Id       int    `json:"id,omitempty"`
}

type AuthResponse struct {
	Token string         `json:"token"`
	User  UserPublicData `json:"user"`
}
