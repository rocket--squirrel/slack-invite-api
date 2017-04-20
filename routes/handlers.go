package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/trickierstinky/slack-invite-api/data"
	"github.com/trickierstinky/slack-invite-api/services"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func PostInvite(w http.ResponseWriter, r *http.Request) {
	var inviteResponse data.InviteResponse
	r.ParseForm()
	//Insert Code to call Slack
	form := []byte(strings.Join(r.Form["payload"], ""))
	if err := json.Unmarshal(form, &inviteResponse); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	fmt.Println(inviteResponse.Actions[0].Value)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(200)

	var message string
	response := strings.Split(inviteResponse.Actions[0].Value, ";")
	invite_id, _ := strconv.Atoi(response[1])
	invite := data.FetchInvite(invite_id)
	if invite != (data.Invite{}) {
		if response[0] == "yes" {
			message = fmt.Sprintf("We've Invited %s", invite.Name)
			services.SendSlackInviteRequest(invite.Email)
		} else {
			message = fmt.Sprintf("We've Not Invited %s", invite.Name)
		}
		fullMessage := fmt.Sprintf("{\"replace_original\": true,\"text\": \"Thanks for the update, %s\"}", message)
		fmt.Fprint(w, fullMessage)

		data.DeleteInvite(invite)
	} else {
		fullMessage := fmt.Sprintf("{\"replace_original\": true,\"text\": \"Problem Updating\"}")
		fmt.Fprint(w, fullMessage)
	}
}

func PostIndex(w http.ResponseWriter, r *http.Request) {
	var invite data.Invite
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &invite); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	//Insert Code to call Slack
	invite = data.CreateInvite(invite)

	services.PostSlackInviteRequest(invite)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)

	if err := json.NewEncoder(w).Encode(invite); err != nil {
		panic(err)
	}
}
