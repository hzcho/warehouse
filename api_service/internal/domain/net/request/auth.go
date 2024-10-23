package request

type SignUp struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type SignIn struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RefreshToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}