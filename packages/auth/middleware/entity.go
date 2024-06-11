package middlewareAuth

type JwtPayload struct {
	Token string `json:"Token"`
}

type Request struct {
	Token       string `json:"Token"`
	AccessToken string
}
