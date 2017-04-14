package data

type Invite struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Description string `json:"description"`
}

type Invites []Invite

type InviteResponse struct {
	Actions []InviteResponseAction
}

type InviteResponseAction struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Value string `json:"value"`
}
