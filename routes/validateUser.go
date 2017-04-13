package routes

import (
	"github.com/trickierstinky/slack-invite-api/data"
)

//replace for something stronger
func validateUser(username, password string) bool {
	if username == env.fetch("username") && password == env.fetch("password") {
		return true
	}
	return false
}
