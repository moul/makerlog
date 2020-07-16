package makerlog

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/url"

	"moul.io/makerlog/makerlogtypes"
	"moul.io/roundtripper"
)

type Client struct {
	http     http.Client
	hasToken bool
}

func New(token string) *Client {
	client := Client{}

	transport := roundtripper.Transport{
		ExtraHeader: http.Header{"Content-Type": []string{"application/json"}},
	}
	if token != "" {
		transport.ExtraHeader["Authorization"] = []string{"Token " + token}
		client.hasToken = true
	}
	client.http = http.Client{Transport: &transport}

	return &client
}

func Login(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", errors.New("missing username or password")
	}

	formData := url.Values{
		"username": {username},
		"password": {password},
	}

	// FIXME: use context
	resp, err := http.PostForm("https://api.getmakerlog.com/api-token-auth/", formData)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var reply struct {
		Token string `json:"token"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return "", err
	}

	return reply.Token, nil
}

func (c *Client) RawNotificationsList(ctx context.Context) (*makerlogtypes.NotificationsListReply, error) {
	resp, err := c.http.Get("https://api.getmakerlog.com/notifications/")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var reply makerlogtypes.NotificationsListReply
	if err := json.NewDecoder(resp.Body).Decode(&reply); err != nil {
		return nil, err
	}

	return &reply, nil
}
