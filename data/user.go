package data

import (
	"github.com/trickierstinky/slack-invite-api/config"
)

//replace for something stronger
func ValidateUser(username, password string) bool {
	if username == config.Env("username") && password == config.Env("password") {
		return true
	}
	return false
}
