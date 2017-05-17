package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/nlopes/slack"
	"github.com/trickierstinky/slack-invite-api/config"
	"github.com/trickierstinky/slack-invite-api/data"
)

//https://api.slack.com/custom-integrations/legacy-tokens

func SendSlackInviteRequest(email string, name string) error {

	var Url *url.URL
	Url, err := url.Parse("https://slack.com")
	if err != nil {
		panic("boom")
	}

	Url.Path += "/api/users.admin.invite"
	parameters := url.Values{}
	parameters.Add("token", config.Env("slack_invite_token"))
	parameters.Add("email", email)
	Url.RawQuery = parameters.Encode()

	_, err2 := http.Get(Url.String())
	if err2 != nil {
		fmt.Printf("%s (%s)\n", err2, email)
	}
	return err2
}

func PostSlackInviteRequest(invite data.Invite) {
	api := slack.New(config.Env("slack_token"))
	api.SetDebug(false)

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
				Value: fmt.Sprintf("yes;%d", int(invite.ID)),
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
				Value: fmt.Sprintf("no;%d", int(invite.ID)),
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
	channelID, timestamp, err := api.PostMessage(config.Env("slack_channel_id"), "New Invite Request", params)
	if err != nil {
		fmt.Printf("%s (%s)\n", err, timestamp)
		return
	}
	log.Printf("Message successfully sent to channel %s at %s", channelID, timestamp)
}
