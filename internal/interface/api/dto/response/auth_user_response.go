package response

type AuthUserResponse struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}
