package routes

//replace for something stronger
func validateUser(username, password string) bool {
	if username == "test" && password == "test" {
		return true
	}
	return false
}
