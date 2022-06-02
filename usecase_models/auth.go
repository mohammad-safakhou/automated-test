package usecase_models

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

type AccountAuth struct {

}
