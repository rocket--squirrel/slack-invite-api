package data

var invites Invites


func RepoCreateInvite(i Invite) Invite {
	invites = append(invites, i)
	return i
}
