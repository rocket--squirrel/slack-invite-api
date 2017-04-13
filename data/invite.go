package data

type Invite struct {
	Name				string	`json:"name"`
	Email				string	`json:"email"`
	Description	string	`json:"description"`
}

type Invites []Invite
