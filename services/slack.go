package services

import (
	"fmt"

	"github.com/nlopes/slack"
	"github.com/trickierstinky/slack-invite-api/config"
	"github.com/trickierstinky/slack-invite-api/data"
)

//https://slack.com/oauth/authorize?
//&client_id=CLIENTID
//&team=TEAMID&install_redirect=oauth&scope=client

func SendSlackInviteRequest(email string) bool {
	api := slack.New(config.Env("slack_token"))
	api.SetDebug(true)
	fmt.Printf("%s (%s)\n", config.Env("slack_token"), config.Env("slack_team"))

	err := api.InviteToTeam(config.Env("slack_team"), "test", "test", email)
	if err != nil {
		fmt.Printf("%s (%s)\n", err, email)
		return false
	}
	return true
}

func PostSlackInviteRequest(invite data.Invite) {
	api := slack.New(config.Env("slack_token"))
	api.SetDebug(true)

	params := slack.PostMessageParameters{}
	attachment := slack.Attachment{
		Text: fmt.Sprintf("We've had a new Invite Request from *%s* `%s` ```%s``` ", invite.Name, invite.Email, invite.Description),
		MarkdownIn: []string{
			"text",
			"fields",
		},
		CallbackID: "RS_INVITE",
		Color:      "#36a64f",
		AuthorLink: "Rocket The Squirrel",
		AuthorIcon: "https://avatars.slack-edge.com/2016-09-30/85928175318_c13100913436073bc926_48.jpg",
		Actions: []slack.AttachmentAction{
			slack.AttachmentAction{
				Name:  "response",
				Text:  "Approve",
				Type:  "button",
				Value: fmt.Sprintf("yes;%s", invite.Email),
				Confirm: &slack.ConfirmationField{
					Title:       "Are you sure?",
					Text:        fmt.Sprintf("This will automatically, send out an email to %s inviting them to the group.", invite.Email),
					OkText:      "Yes",
					DismissText: "No",
				},
			},
			slack.AttachmentAction{
				Name:  "response",
				Text:  "Reject",
				Type:  "button",
				Value: fmt.Sprintf("no;%s", invite.Email),
				Confirm: &slack.ConfirmationField{
					Title:       "Are you sure?",
					Text:        "This will reject the invite request",
					OkText:      "Yes",
					DismissText: "No",
				},
			},
		},
	}

	params.Attachments = []slack.Attachment{attachment}
	// fmt.Println(params)
	channelID, timestamp, err := api.PostMessage(config.Env("channel_id"), "New Invite Request", params)
	if err != nil {
		fmt.Printf("%s (%s)\n", err, timestamp)
		return
	}
	fmt.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
