package data

//replace for something stronger
func ValidateUser(username, password string) bool {
	if username == Env("username") && password == Env("password") {
		return true
	}
	return false
}
