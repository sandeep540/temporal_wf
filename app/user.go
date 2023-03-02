package app

const SignupUserQueue = "SIGNUP_USER_QUEUE"

// @@@SNIPEND

type UserDetails struct {
	name    string
	email   string
	country string
}

type Init struct {
	sessionID string `json:"sessionId"`
	JWT       string `json:"jwt"`
	customer  string `json:"customer"`
}
