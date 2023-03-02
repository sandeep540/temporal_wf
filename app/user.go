package app

const SignupUserQueue = "SIGNUP_USER_QUEUE"

// @@@SNIPEND

type Init struct {
	sessionId string `json:"sessionId"`
	JWT       string `json:"jwt"`
	customer  string `json:"customer"`
}
